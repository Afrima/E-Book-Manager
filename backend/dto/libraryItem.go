package dto

type LibraryItem struct {
	Cover     *string `json:"cover"`
	Title     string  `json:"title"`
	ItemType  string  `json:"itemType"`
	BookCount uint    `json:"bookCount"`
	Id        uint    `json:"id"`
}
