package db

type Author2Book struct {
	Author_id uint   `gorm:"primaryKey"`
	Book_id   uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null; default:aut"`
}
