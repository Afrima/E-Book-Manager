package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"fmt"
	"strings"
)

func GetCollection(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Metafield, e *ParseError) uint {
	var collections = make([]string, 0)
	for _, titleMeta := range bookFile.Opf.Metadata.Title {
		if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "collection" {
			collections = append(collections, strings.TrimSpace(titleMeta.Data))
		}
	}
	for _, metafield := range bookFile.Opf.Metadata.Meta {
		if strings.HasSuffix(metafield.Name, "series") {
			var collectionName = strings.TrimSpace(metafield.Data)
			if len(collectionName) == 0 {
				collectionName = strings.TrimSpace(metafield.Content)
			}
			collections = append(collections, collectionName)
		}
	}
	if len(collections) > 1 {
		e.Collection = "to many collections"
		return persistCol(collections[0])
	}
	if len(collections) == 0 {
		e.Collection = "not found!"
		return 0
	}
	return persistCol(collections[0])
}

func persistCol(name string) uint {
	fmt.Println(name)
	var collection = book.GetCollectionByName(name)
	if collection.Name == "" {
		collection.Name = name
		collection.Persist()
	}
	return collection.ID
}

/*
if bookEntity.CollectionId == 0 {
		for _, metafield := range bookFile.Opf.Metadata.Meta {
			if strings.HasSuffix(metafield.Name, "series") {
				var collectionName = strings.TrimSpace(metafield.Data)
				if len(collectionName) == 0 {
					collectionName = strings.TrimSpace(metafield.Content)
				}
				var collection = book.GetCollectionByName(collectionName)
				if collection.Name == "" {
					collection.Name = collectionName
					collection.Persist()
				}
				bookEntity.CollectionId = collection.ID
				break
			}
		}
	}
*/
