package DAL

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func connectmgo() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		log.Print(err)
	}
	return session

}

func GetAllBooks() []Book {
	var books []Book
	session := connectmgo()
	defer session.Close()

	//to connect to a collection
	c := session.DB("weinardb").C("books") //will get collection object
	err := c.Find(nil).All(&books)

	log.Print(books)

	if err != nil {
		log.Print(err)
	}
	return books
}

func GetBookByIsbn(isbn string) Book {

	session := connectmgo()
	defer session.Close()
	c := session.DB("weinardb").C("books")

	var book Book
	err := c.Find(bson.M{"isbn": isbn}).One(&book)
	if err != nil {
		log.Print(err)
	}
	return book
}

func AddBook(book Book) {
	session := connectmgo()
	defer session.Close()
	c := session.DB("weinardb").C("books")
	err := c.Insert(book)
	if err != nil {
		log.Print(err)
	}
}

func DeleteBookmgo(isbn string) {
	log.Print("in deletebookmgo")
	session := connectmgo()
	defer session.Close()
	c := session.DB("weinardb").C("books")
	err := c.Remove(bson.M{"isbn": isbn})
	if err != nil {
		log.Print(err)
	}
}

func UpdateBookMongo(book Book) {
	log.Print("in updatebookmgo")
	session := connectmgo()
	defer session.Close()
	c := session.DB("weinardb").C("books")
	err := c.Update(bson.M{"isbn": book.Isbn}, &book)
	if err != nil {
		log.Print(err)
	}
}
