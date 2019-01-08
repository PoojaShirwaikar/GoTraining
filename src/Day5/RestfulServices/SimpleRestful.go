package main

import (
	"Day5/RestfulServices/DAL"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "GET BOOKS IS CALLED..")
	//books := DAL.GetTempBooks()

	//mysql call
	//books := DAL.GetBooks()

	//mongodb call
	log.Print("connected to mongodb...")
	books := DAL.GetAllBooks()
	log.Print(books)
	//converting books slice to json object  (return json coz other languages may not understand goland slice)
	//can use json.Marshal also
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "GET BOOK by id IS CALLED..")

	//"/books/{id}"

	params := mux.Vars(r)
	id := params["id"]
	value, _ := strconv.Atoi(id)
	book := DAL.GetBookById(value)
	json.NewEncoder(w).Encode(book)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "post BOOK IS CALLED..")

	//request will have book object in json format
	//convert json object to book object

	var book DAL.Book

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&book)
	log.Print(&book)
	DAL.PostBook(book)
	log.Print(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "delete BOOK IS CALLED..")
	params := mux.Vars(r)
	id := params["id"]
	value, _ := strconv.Atoi(id)
	DAL.Delete(value)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "update BOOK IS CALLED..")
	var book DAL.Book

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&book)
	log.Print(&book)
	//sql call
	//	DAL.UpdateBook(book)

	//mongodb call
	DAL.UpdateBookMongo(book)
	log.Print(book)
}

func GetBookByIsbn(w http.ResponseWriter, r *http.Request) {
	log.Print("in getbookbyisbn")
	parms := mux.Vars(r)
	isbn := parms["isbn"]
	book := DAL.GetBookByIsbn(isbn)

	json.NewEncoder(w).Encode(book)
}

func PostMongoDB(w http.ResponseWriter, r *http.Request) {

	var book DAL.Book

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&book)
	log.Print(&book)
	DAL.AddBook(book)
	log.Print(book)
}

func DeleteBookmgodb(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "delete BOOK IS CALLED..")
	params := mux.Vars(r)
	isbn := params["isbn"]

	DAL.DeleteBookmgo(isbn)
}

func main() {
	//using MUX package define restful endpoint
	//http verb
	//request url
	//decide which handler to call

	var router = mux.NewRouter()
	//if url is http://localhost:8080/books  :GET
	//configured different routes
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	router.HandleFunc("/books", PostBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")

	router.HandleFunc("/books/isbn/{isbn}", GetBookByIsbn).Methods("GET")
	router.HandleFunc("/books/mgo", PostMongoDB).Methods("POST")
	router.HandleFunc("/books/mgo/{isbn}", DeleteBookmgodb).Methods("DELETE")

	router.HandleFunc("/books/isbn/{isbn}", UpdateBook).Methods("PUT")

	//listen to some port
	http.ListenAndServe(":8080", router)
}
