package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//import "github.com/go-sql-driver/mysql"

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/sample")

	// if there is an error opening the connection, handle it
	if err != nil {
		//log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT PersonID, FirstName FROM Persons")

	fmt.Println(results)

	for results.Next() {
		//var uid int
		var PersonID int
		var FirstName string
		err = results.Scan(&PersonID, &FirstName)
		//checkErr(err)
		fmt.Println(PersonID)
		fmt.Println(FirstName)

	}
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// for results.Next() {
	// 	var tag Tag
	// 	// for each row, scan the result into our tag composite object
	// 	err = results.Scan(&tag.PersonID, &tag.FirstName)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	//             // and then print out the tag's Name attribute
	// 	log.Printf(tag.FirstName)
	// }

}
