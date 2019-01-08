package DAL

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//used blank identifier only init() getting called
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Print(err)
	}
	return db
}

func GetBooks() []Book {
	var db = connect()

	//close connection before returning from GetBooks
	defer db.Close()

	//different methods of Quering
	//1. return multiple records
	//2.return single record

	//return collection of rows
	rows, err := db.Query("Select ID,name,isbn,author,pages from Books")
	if err != nil {
		log.Print(err)
	}

	var books []Book
	//iterate through rows collection
	var book Book
	for rows.Next() {
		//map records to variable
		err1 := rows.Scan(&book.ID, &book.Name, &book.Isbn, &book.Author, &book.Pages)
		if err1 != nil {
			log.Print(err1)
		}
		//append book object to book slice
		books = append(books, book)
	}
	return books
}

func GetBookById(id int) Book {
	var db = connect()

	defer db.Close()

	var book Book
	row := db.QueryRow("Select ID,name,isbn,author,pages from Books where ID=?;", id)
	err2 := row.Scan(&book.ID, &book.Name, &book.Isbn, &book.Author, &book.Pages)
	if err2 != nil {
		log.Print(err2)
	}
	return book
}

func PostBook(book Book) {
	var db = connect()
	defer db.Close()

	//can execute same prepared stmt multiple times with same instance
	stmt, _ := db.Prepare("insert into Books(name, isbn,author,pages) values(?,?,?,?);")
	_, err := stmt.Exec(book.Name, book.Isbn, book.Author, book.Pages)

	if err != nil {
		log.Print(err)

	}

}

func UpdateBook(book Book) {
	var db = connect()
	defer db.Close()

	stmt, _ := db.Prepare("Update Books set name=?, isbn=?,author=?,pages=? where id=?;")
	_, err := stmt.Exec(book.Name, book.Isbn, book.Author, book.Pages, book.ID)

	if err != nil {
		log.Print(err)

	}

}

func Delete(id int) {
	var db = connect()
	defer db.Close()

	stmt, err := db.Prepare("Delete from Books where ID=?;")

	if err != nil {
		log.Print(err)
	}

	stmt.Exec(id)
}

type Book struct {
	ID   int
	Name string `json:"bookname,omitempty"`
	Isbn string `json:"isbn"`
	//isbn  (not capitalized so wont be exported. encoded or decoded)
	Author string `json:"-"` //should be accesible within ackage bt ignored for encoding
	Pages  int    `json:"pages"`
}

func GetTempBooks() []Book {
	var books []Book
	books = append(books, Book{1, "GoLANG", "isbn123", "rob", 100})
	books = append(books, Book{2, "MySQL", "isbn123", "rob", 200})
	books = append(books, Book{3, "MongoDB", "isbn123", "rob", 300})

	return books

}
