package library

import (
	"context"
	"errors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/infrastructure"
	"strings"
)

var (
	bookRepo   *book.Repository
	authorRepo *author.Repository
)

var (
	ErrBookOutOfStock = errors.New("there is not enough stock to sell this book in demanded amount")
)

// Init initializes DB, repositories and insert sample data if DB is empty
func Init() {
	db := infrastructure.NewPostgresDB()
	bookRepo = book.NewBookRepository(db)
	authorRepo = author.NewAuthorRepository(db)

	err := authorRepo.Migration()
	if err != nil {
		panic(err)
	}
	err = bookRepo.Migration()
	if err != nil {
		panic(err)
	}

	insertSampleAuthors()
	insertSampleBooks()
}

// List service layer of list operation
// it returns all books not deleted
func List(ctx context.Context) ([]book.Book, error) {
	c, e := make(chan interface{}, 1), make(chan error, 1)
	go func() {
		books, err := bookRepo.FindAllWithAuthor(ctx)
		if err != nil {
			e <- err
			return
		}
		c <- books
	}()

	books, err := waitResponse(ctx, c, e)
	if err != nil {
		return nil, err
	}
	return books.([]book.Book), nil
}

// Search service layer of search operation
// it finds books not deleted and meet criteria
func Search(ctx context.Context, text []string) ([]book.Book, error) {
	c, e := make(chan interface{}, 1), make(chan error, 1)
	go func() {
		// join texts with space and search by it
		books, err := bookRepo.Search(ctx, strings.Join(text, " "))
		if err != nil {
			e <- err
			return
		}
		c <- books
	}()

	books, err := waitResponse(ctx, c, e)
	if err != nil {
		return nil, err
	}
	return books.([]book.Book), nil
}

// Buy service layer of buy operation
// it decreases stock amount of book with given id by quantity
func Buy(ctx context.Context, id, quantity int) (book.Book, error) {
	c, e := make(chan interface{}, 1), make(chan error, 1)
	go func() {
		b, err := bookRepo.FindById(ctx, id)
		if err != nil {
			e <- err
			return
		}
		if b.StockAmount >= quantity {
			b.StockAmount -= quantity
			err = bookRepo.Update(ctx, b)
			if err != nil {
				e <- err
				return
			}
			c <- b
			return
		}
		e <- ErrBookOutOfStock
	}()

	b, err := waitResponse(ctx, c, e)
	if err != nil {
		return book.Book{}, err
	}
	return b.(book.Book), nil
}

// Delete service layer of delete operation
// deletes book with given id if it is nor already deleted
func Delete(ctx context.Context, id int) (book.Book, error) {
	c, e := make(chan interface{}, 1), make(chan error, 1)
	go func() {
		b, err := bookRepo.FindById(ctx, id)
		if err != nil {
			e <- err
			return
		}
		err = bookRepo.Delete(ctx, b)
		if err != nil {
			e <- err
			return
		}
		c <- b
	}()

	b, err := waitResponse(ctx, c, e)
	if err != nil {
		return book.Book{}, err
	}
	return b.(book.Book), nil
}

// Clear it operates hard delete for all records
func Clear() error {
	err := bookRepo.Clear()
	if err != nil {
		return err
	}
	err = authorRepo.Clear()
	if err != nil {
		return err
	}
	return nil
}

// waitResponse waits response from context or channel and returns
func waitResponse(ctx context.Context, c chan interface{}, e chan error) (interface{}, error) {
	select {
	case m := <-c:
		return m, nil
	case err := <-e:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
