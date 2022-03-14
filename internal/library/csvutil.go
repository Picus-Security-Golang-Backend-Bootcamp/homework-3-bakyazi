package library

import (
	"context"
	"encoding/csv"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/book"
	"log"
	"os"
	"strconv"
)

// readCsvFile opens and reads CSV file name @filename
func readCsvFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("cannot open sample csv file")
		panic(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		log.Println("cannot read sample csv file")
		panic(err)
	}
	return data[1:]
}

// csvToBookSlice read a csv file and converts it to []book.Book
func csvToBookSlice(filename string) []book.Book {
	data := readCsvFile(filename)
	books := make([]book.Book, len(data))
	for i, row := range data {
		id, _ := strconv.Atoi(row[0])
		authorId, _ := strconv.Atoi(row[2])
		pageCount, _ := strconv.Atoi(row[5])
		price, _ := strconv.Atoi(row[6])
		stockAmount, _ := strconv.Atoi(row[7])
		books[i] = book.Book{
			ID:          id,          // index 0
			Name:        row[1],      // index 1
			AuthorID:    authorId,    // index 2
			StockCode:   row[3],      // index 3
			ISBN:        row[4],      // index 4
			PageCount:   pageCount,   // index 5
			Price:       price,       // index 6
			StockAmount: stockAmount, // index 7
		}
	}
	return books
}

// csvToAuthorSlice read a csv file and converts it to []author.Author
func csvToAuthorSlice(filename string) []author.Author {
	data := readCsvFile(filename)
	authors := make([]author.Author, len(data))
	for i, row := range data {
		id, _ := strconv.Atoi(row[0])
		authors[i] = author.Author{
			ID:   id,     // index 0
			Name: row[1], // index 1
		}
	}
	return authors
}

func insertSampleAuthors() {
	savedAuthors, err := authorRepo.FindAll(context.Background(), true)
	if err != nil {
		log.Printf("cannot get all author data from DB, %v\n", err)
	}
	if len(savedAuthors) == 0 {
		authors := csvToAuthorSlice("resources/author.csv")
		err := authorRepo.BulkInsert(authors)
		if err != nil {
			log.Printf("cannot insert sample author data, %v\n", err)
		}
		return
	}
	log.Printf("there are already %d authors record in DB\n", len(savedAuthors))

}

func insertSampleBooks() {
	savedBooks, err := bookRepo.FindAll(context.Background(), true)
	if err != nil {
		log.Printf("cannot get all book data from DB, %v\n", err)
	}
	if len(savedBooks) == 0 {
		books := csvToBookSlice("resources/book.csv")
		err := bookRepo.BulkInsert(books)
		if err != nil {
			log.Printf("cannot insert sample book data, %v\n", err)
		}
		return
	}
	log.Printf("there are already %d books record in DB\n", len(savedBooks))
}
