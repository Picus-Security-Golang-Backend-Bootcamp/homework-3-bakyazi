# homework-3-bakyazi

Author: Berkay Akyazi

## build

```shell
$ go mod download
$ go build -o bin/library cmd/hw3/main.go
```

## usage
```shell
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
```

## environment variables

There should be two files (`.env` or `.env.local`) at root directory of the project

`.env` file must exist and `PATIKA_ENV_PROFILE` must be set. If this parameter is `PROD` then project uses values from `.env` otherwise it overwrites values with `.env.local`


## example usages

### list
```shell
$ ./bin/library list 

RESULT:
        
1. In Search of Lost Time (ID=1) (ISBN=396-85-54496-53-8) (Price=$132) (StockAmount=133) | [Author (ID=108) (Name=Marcel Proust)]
2. Ulysses (ID=2) (ISBN=307-20-73436-29-8) (Price=$62) (StockAmount=22) | [Author (ID=133) (Name=James Joyce)]
3. Don Quixote (ID=3) (ISBN=115-73-93054-57-5) (Price=$14) (StockAmount=179) | [Author (ID=52) (Name=Miguel de Cervantes)]
4. One Hundred Years of Solitude (ID=4) (ISBN=507-91-50245-10-5) (Price=$112) (StockAmount=25) | [Author (ID=39) (Name=Gabriel Garcia Marquez)]
5. The Great Gatsby (ID=5) (ISBN=667-63-23448-43-2) (Price=$37) (StockAmount=174) | [Author (ID=110) (Name=F. Scott Fitzgerald)]
6. Moby Dick (ID=6) (ISBN=875-81-76123-48-8) (Price=$115) (StockAmount=170) | [Author (ID=1) (Name=Herman Melville)]
7. War and Peace (ID=7) (ISBN=638-69-17646-66-6) (Price=$76) (StockAmount=64) | [Author (ID=41) (Name=Leo Tolstoy)]
8. Hamlet (ID=8) (ISBN=274-93-91976-48-4) (Price=$96) (StockAmount=145) | [Author (ID=147) (Name=William Shakespeare)]
.
.
198. Atlas Shrugged (ID=197) (ISBN=906-27-65315-31-5) (Price=$64) (StockAmount=171) | [Author (ID=111) (Name=Ayn Rand)]
199. Winnie the Pooh (ID=199) (ISBN=692-58-95930-93-6) (Price=$49) (StockAmount=185) | [Author (ID=117) (Name=A. A Milne)]
200. A Doll's House (ID=200) (ISBN=326-19-51292-98-1) (Price=$29) (StockAmount=177) | [Author (ID=142) (Name=Henrik Ibsen)]
```

### search

search (not found)
```shell
$ ./bin/library search gfsdfjksdhfjksdfsdfjsd
2022/03/14 21:18:56 there are already 147 authors record in DB
2022/03/14 21:18:56 there are already 200 books record in DB

ERROR:
        Not found any book with respect to search criteria
```


search by book name
```shell
$ ./bin/library search great
2022/03/14 21:18:04 there are already 147 authors record in DB
2022/03/14 21:18:04 there are already 200 books record in DB

RESULT:
        
1. The Great Gatsby (ID=5) (ISBN=667-63-23448-43-2) (Price=$37) (StockAmount=174) | [Author (ID=110) (Name=F. Scott Fitzgerald)]
2. Great Expectations (ID=26) (ISBN=980-99-47637-75-2) (Price=$57) (StockAmount=172) | [Author (ID=140) (Name=Charles Dickens)]
```

search by author
```shell
$ ./bin/library search james                 
2022/03/14 21:19:30 there are already 147 authors record in DB
2022/03/14 21:19:30 there are already 200 books record in DB

RESULT:
        
1. Ulysses (ID=2) (ISBN=307-20-73436-29-8) (Price=$62) (StockAmount=22) | [Author (ID=133) (Name=James Joyce)]
2. A Portrait of the Artist as a Young Man (ID=47) (ISBN=699-78-62934-38-6) (Price=$54) (StockAmount=105) | [Author (ID=133) (Name=James Joyce)]
3. The Portrait of a Lady (ID=59) (ISBN=380-97-96631-77-5) (Price=$125) (StockAmount=163) | [Author (ID=76) (Name=Henry James)]
4. The Ambassadors (ID=156) (ISBN=479-89-69823-71-6) (Price=$108) (StockAmount=179) | [Author (ID=76) (Name=Henry James)]
```

search by ISBN
```shell
$ ./bin/library search 487-56-86624-25-9     
2022/03/14 21:20:12 there are already 147 authors record in DB
2022/03/14 21:20:12 there are already 200 books record in DB

RESULT:
        
1. The Hobbit (ID=180) (ISBN=487-56-86624-25-9) (Price=$98) (StockAmount=32) | [Author (ID=14) (Name=J. R. R. Tolkien)]
```

### buy
successful buy
```shell
$ ./bin/library buy 1 10                
2022/03/14 21:21:45 there are already 147 authors record in DB
2022/03/14 21:21:45 there are already 200 books record in DB

RESULT:
        You bought 10 from Book[ID=1] successfully! There are 123 of this book left
```
failed buy (exceed stock amount)
```shell
$ ./bin/library buy 1 124
2022/03/14 21:22:07 there are already 147 authors record in DB
2022/03/14 21:22:07 there are already 200 books record in DB

ERROR:
        error occured during buy operation, 
         - there is not enough stock to sell this book in demanded amount
```
failed buy (try to buy non-exist book)
```shell
$ ./bin/library buy 599 1  
2022/03/14 21:22:28 there are already 147 authors record in DB
2022/03/14 21:22:28 there are already 200 books record in DB

ERROR:
        error occured during buy operation, 
         - record not found
```
### delete

successful delete
```shell
$ ./bin/library delete 1  
2022/03/14 21:23:40 there are already 147 authors record in DB
2022/03/14 21:23:40 there are already 200 books record in DB

RESULT:
        book[ID=1] is successfully deleted
```

failed delete (already deleted)
```shell
$ ./bin/library delete 1
2022/03/14 21:23:54 there are already 147 authors record in DB
2022/03/14 21:23:54 there are already 200 books record in DB

ERROR:
        error occured during delete operation,
         - record not found
```

failed delete (non exist book)
```shell
$ ./bin/library delete 1000
2022/03/14 21:24:08 there are already 147 authors record in DB
2022/03/14 21:24:08 there are already 200 books record in DB

ERROR:
        error occured during delete operation,
         - record not found
```
### clear
```shell
$ ./bin/library clear                    

RESULT:
        all records successfully deleted
```