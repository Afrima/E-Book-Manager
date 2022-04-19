package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"os"
)

type LibraryItem struct {
	Cover     string
	Title     string
	ItemType  string
	BookCount uint
	Id        uint
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	cover, _ := os.ReadFile(p.Cover)
	return dto.LibraryItem{
		Cover:     cover,
		Title:     p.Title,
		ItemType:  p.ItemType,
		BookCount: p.BookCount,
		Id:        p.Id,
	}
}

func GetAllLibraryItems(page int) []LibraryItem {
	var libraryItems = make([]LibraryItem, 0)
	db.GetDbConnection().Offset(db.SetPage(page)).Limit(db.Limit).Table("books").Select(" COALESCE(collections.id, books.id) as Id, " +
		" books.cover as Cover, COALESCE(collections.title, books.title) AS Title, " +
		" CASE WHEN collections.title IS NOT NULL THEN 'collection' " +
		" ELSE 'book' END AS ItemType, COUNT(*) AS BookCount " +
		"").Joins("left join collections on books.collection_id = collections.id" +
		"").Group("COALESCE(collections.title, books.title)" +
		"").Scan(&libraryItems)
	return libraryItems
}

func GetLibraryItemByCollectionId(id uint64) LibraryItem {
	var libraryItem = LibraryItem{}
	db.GetDbConnection().Table("books").Select(" COALESCE(collections.id, books.id) as Id, "+
		" books.cover as Cover, COALESCE(collections.title, books.title) AS Title, "+
		" CASE WHEN collections.title IS NOT NULL THEN 'collection' "+
		" ELSE 'book' END AS ItemType, COUNT(*) AS BookCount "+
		"").Joins("left join collections on books.collection_id = collections.id"+
		"").Group("COALESCE(collections.title, books.title)"+
		"").Find(&libraryItem, "collections.id = ?", id)
	return libraryItem
}
