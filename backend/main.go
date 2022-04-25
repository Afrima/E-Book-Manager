package main

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/dto"
	"e-book-manager/epub"
	"e-book-manager/parser"
	"errors"
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func uploadFile(fileHeader *multipart.FileHeader) (*book.Book, error) {
	bookFile, err := epub.Open("upload/ebooks/" + fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	defer bookFile.Close()
	return createBookEntity(bookFile, "upload/ebooks/"+fileHeader.Filename)
}

// todo error handling?
func createBookEntity(bookFile *epub.Book, path string) (*book.Book, error) {
	var coverId = ""
	var metaIdMap = make(map[string]map[string]epub.Metafield)
	for _, meta := range bookFile.Opf.Metadata.Meta {
		if meta.Name == "cover" {
			coverId = meta.Content
		} else if meta.Refines != "" {
			if metaIdMap[meta.Refines] == nil {
				metaIdMap[meta.Refines] = make(map[string]epub.Metafield)
			}
			metaIdMap[meta.Refines][meta.Property] = meta
		}
	}
	bookEntity := book.Book{}
	bookEntity.Title = parser.GetTitle(bookFile, metaIdMap)
	if bookEntity.Title == "" {
		return nil, errors.New("no title found")
	}
	bookEntity.Authors = parser.GetAuthor(bookFile, metaIdMap)
	var date, err = parser.GetDate(bookFile)
	if err == nil {
		bookEntity.Published = *date
	}
	bookEntity.Publisher, _ = parser.GetPublisher(bookFile)
	bookEntity.Language, _ = parser.GetLanguage(bookFile)
	bookEntity.Cover, _ = parser.GetCover(coverId, bookFile, bookEntity.Title)
	bookEntity.Subjects = parser.GetSubject(bookFile)
	bookEntity.Book = path
	bookEntity.CollectionIndex = parser.GetCollectionIndex(bookFile)
	bookEntity.CollectionId = parser.GetCollection(bookFile, metaIdMap, bookEntity.Cover)
	bookEntity.Persist()
	return &bookEntity, nil
}

func setupRoutes() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	username, userNameSetted := os.LookupEnv("user")
	password, passwordSetted := os.LookupEnv("password")
	var auth *gin.RouterGroup = nil
	if userNameSetted && passwordSetted {
		auth = r.Group("/", gin.BasicAuth(gin.Accounts{
			username: password,
		},
		))
	} else {
		auth = r.Group("/")
	}
	r.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		static.Serve("/", static.LocalFile("./bundles", true))(c)
	})
	r.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		c.File("./bundles/index.html")
	})
	auth.POST("/api/upload/multi", func(c *gin.Context) {
		files, _ := c.MultipartForm()
		fileErrors := ""
		for _, fileHeader := range files.File["myFiles"] {
			if fileHeader.Header.Get("Content-Type") != "application/epub+zip" {
				fileErrors += "Error: Book " + fileHeader.Filename + ": is not in epub format\n"
				continue
			}
			err := c.SaveUploadedFile(fileHeader, "upload/ebooks/"+fileHeader.Filename)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			_, err = uploadFile(fileHeader)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
		}
		if len(fileErrors) > 0 {
			c.String(400, fileErrors)
		} else {
			c.JSON(200, "Done")
		}
	})
	auth.GET("/api/library/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		entity := book.GetLibraryItemByCollectionId(id)
		c.JSON(200, entity.ToDto())
	})
	auth.GET("/api/book", func(c *gin.Context) {
		queryParam, exist := c.GetQuery("q")
		if !exist {
			c.String(400, "query param q expected")
		}
		pageQuery, exist := c.GetQuery("page")
		if !exist {
			pageQuery = "1"
		}
		page, err := strconv.ParseUint(pageQuery, 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		search := strings.Split(queryParam, " ")

		var books = book.SearchBooks(search, int(page))
		bookDtos := make([]dto.Book, len(books))
		for i, b := range books {
			bookDtos[i] = b.ToDto()
		}
		c.JSON(200, bookDtos)
	})
	auth.GET("/api/book/:title", func(c *gin.Context) {
		title := c.Param("title")
		entity := book.GetBookByTitle(title)
		c.JSON(200, entity.ToDto())
	})
	auth.GET("/api/collection", func(c *gin.Context) {
		title := c.Query("title")
		byName := book.GetCollectionByName(title)
		c.JSON(200, byName.ToDto())
	})
	auth.GET("/api/collection/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		byName := book.GetCollectionById(id)
		c.JSON(200, byName.ToDto())
	})
	auth.GET("/api/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookEntity := book.GetBookById(id)
		b, err := os.ReadFile(bookEntity.Book)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.Data(200, "application/epub+zip", b)
	})
	auth.GET("/api/all", func(c *gin.Context) {
		pageQuery, exist := c.GetQuery("page")
		if !exist {
			pageQuery = "1"
		}
		page, err := strconv.ParseUint(pageQuery, 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		var libraryItems = book.GetAllLibraryItems(int(page))

		var libraryItemDtos = make([]dto.LibraryItem, len(libraryItems))
		for i, libraryItem := range libraryItems {
			libraryItemDtos[i] = libraryItem.ToDto()
		}
		c.JSON(200, libraryItemDtos)
	})
	auth.GET("/api/rescan", func(c *gin.Context) {
		var files []string

		root := "upload/ebooks/"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			bookFile, err := epub.Open(file)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			defer bookFile.Close()
			createBookEntity(bookFile, file)
		}
	})
	if err := r.Run(":8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	err := os.MkdirAll("upload/ebooks/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = os.MkdirAll("upload/covers/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	//todo maybe in a different place?
	dbCon := db.GetDbConnection()
	err = dbCon.AutoMigrate(&book.Book{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Author{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Subject{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Collection{})
	if err != nil {
		panic(err.Error())
	}
	setupRoutes()
}
