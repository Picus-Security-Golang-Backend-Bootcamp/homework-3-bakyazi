package book

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	ID          int `gorm:"primary_key"`
	Name        string
	AuthorID    int
	StockCode   string
	ISBN        string
	PageCount   int
	Price       int
	StockAmount int
	Deleted     gorm.DeletedAt
	Author      author.Author `gorm:"foreignKey:AuthorID"`
}

func (b Book) String() string {
	var authorName = "% not fetched %"
	if b.Author.ID > 0 {
		authorName = b.Author.Name
	}
	return fmt.Sprintf("%s (ID=%d) (ISBN=%s) (Price=$%d) (StockAmount=%d) | [Author (ID=%d) (Name=%s)]",
		b.Name,
		b.ID,
		b.ISBN,
		b.Price,
		b.StockAmount,
		b.AuthorID,
		authorName)
}
