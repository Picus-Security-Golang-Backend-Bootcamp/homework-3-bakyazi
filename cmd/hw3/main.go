package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-bakyazi/internal/library"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	EnvResponseTimeoutKey string = "PATIKE_RESP_TIMEOUT_IN_MSEC"
	DefaultRespTimeout           = time.Millisecond * 5000
	SuccessPrefix                = "\nRESULT:\n\t"
	ErrorPrefix                  = "\nERROR:\n\t"
)

// init loads books slice as a library.BookList
func init() {
	err := loadEnv()
	if err != nil {
		log.Printf("cannot load env,err: %v\n", err)
	}
	library.Init()
}

// loadEnv loads env variables from .env or .env.local
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	// if LOCAL mode enable env variables are overwritten by .env.local
	if os.Getenv("PATIKA_ENV_PROFILE") == "LOCAL" {
		err := godotenv.Overload(".env.local")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	args := os.Args

	if len(args) == 1 {
		printUsage()
		return
	}

	command := args[1]

	switch command {
	case "search":
		searchOperation(args)
	case "list":
		listOperation()
	case "delete":
		deleteOperation(args)
	case "buy":
		buyOperation(args)
	case "clear":
		err := library.Clear()
		if err != nil {
			fmt.Printf("%serror occured in clear operation, %v\n", ErrorPrefix, err)
			return
		}
		fmt.Printf("%sall records successfully deleted\n", SuccessPrefix)
	default:
		printUsage()
	}

}

// listOperation operates list command
func listOperation() {
	ctx, cancel := context.WithTimeout(context.Background(), getRespTimeout())
	defer cancel()

	books, err := library.List(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("%slist operation takes too much time! please increase timeout configuration, %v\n", ErrorPrefix, err)
		} else {
			fmt.Printf("%serror occured in list operation, %v\n", ErrorPrefix, err)
		}
		return
	}
	if len(books) == 0 {
		fmt.Printf("%sThere is no book in library!\n", ErrorPrefix)
		return
	}
	printBooks(books)
	return
}

// searchOperation operates search command
func searchOperation(args []string) {
	if len(args) < 3 {
		printUsage()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), getRespTimeout())
	defer cancel()

	books, err := library.Search(ctx, args[2:])
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("%ssearch operation takes too much time! please increase timeout configuration, %v\n", ErrorPrefix, err)
		} else {
			fmt.Printf("%serror occured in search operation, %v\n", ErrorPrefix, err)
		}
		return
	}
	if len(books) == 0 {
		fmt.Printf("%sNot found any book with respect to search criteria\n", ErrorPrefix)
		return
	}
	printBooks(books)
	return
}

// buyOperation operates buy command
func buyOperation(args []string) {
	if len(args) != 4 {
		printUsage()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("%sgiven id (%s) is not valid! please enter valid integer! \n", ErrorPrefix, args[2])
		return
	}

	amount, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("%sgiven amount (%s) is not valid! please enter valid integer! \n", ErrorPrefix, args[3])
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), getRespTimeout())
	defer cancel()

	item, err := library.Buy(ctx, id, amount)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("%sbuy operation takes too much time! please increase timeout configuration, %v\n", ErrorPrefix, err)
		} else {
			fmt.Printf("%serror occured during buy operation, \n\t - %s\n", ErrorPrefix, err.Error())
		}
		return
	}
	fmt.Printf("%sYou bought %d from Book[ID=%d] successfully! There are %d of this book left\n",
		SuccessPrefix,
		amount,
		item.ID,
		item.StockAmount,
	)
}

// deleteOperation operates delete command
func deleteOperation(args []string) {
	if len(args) != 3 {
		printUsage()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("%sgiven id cannot be converted to integer %s\n", ErrorPrefix, args[2])
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), getRespTimeout())
	defer cancel()

	_, err = library.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("%sdelete operation takes too much time! please increase timeout configuration, %v\n", ErrorPrefix, err)
		} else {
			fmt.Printf("%serror occured during delete operation, \n\t - %s\n", ErrorPrefix, err.Error())
		}
		return
	}
	fmt.Printf("%sbook[ID=%d] is successfully deleted\n", SuccessPrefix, id)
	return
}

//printUsage prints usage
func printUsage() {
	usage := `
### USAGE ###

Commands: 
- search => to search and list books with specified arguments, arguments are searched in books' name, author, isdn, stockCode and ID attributes
	e.g:
		$ ./bin/library search moby dick
- list => to show list of all books
	e.g:
		$ ./bin/library list
- delete => to delete the book by specified ID
	e.g:
		$ ./bin/library delete 5
- buy => to buy the book specified by the ID in the specified amount, first argument is ID of the book and second argument is the amount desired to be bought
	e.g:
		$ ./bin/library buy 5 10
- clear => to operate hard delete for all tables. only recommended for reset DB with sample inputs
	e.g:
		$ ./bin/library clear


### ENVIRONMENT VARIABLES ###

- PATIKA_ENV_PROFILE => Run mode (PROD/LOCAL)
- PATIKA_DB_HOST => Host/IP Address for DB connection (localhost, 127.0.0.1)
- PATIKA_DB_PORT => Port of DB connection
- PATIKA_DB_USER => DB Username  
- PATIKA_DB_PASSWORD => DB Password
- PATIKA_DB_NAME => DB Name
- PATIKE_RESP_TIMEOUT_IN_MSEC => Response Timeout in milliseconds
`
	fmt.Println(usage)
}

// printBooks prints book list prettier
func printBooks(books []book.Book) {
	fmt.Printf("%s\n", SuccessPrefix)
	for i, b := range books {
		fmt.Printf("%d. %s\n", i+1, b)
	}
}

func getRespTimeout() time.Duration {
	val := os.Getenv(EnvResponseTimeoutKey)
	if val == "" {
		return DefaultRespTimeout
	}
	t, err := strconv.Atoi(val)
	if err != nil {
		return DefaultRespTimeout
	}
	return time.Millisecond * time.Duration(t)
}
