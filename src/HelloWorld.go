package main

import (
	"DataAccess"
	"DataAccess/DB"
	dbconfig "DataAccess/DB/Config"
	fileconfig "DataAccess/File/Config"
	"fmt"

	"github.com/golang/example/stringutil"
)

//main func would be entry point
func main() {
	fmt.Println("Hello World")
	DataAccess.ReadData()
	DataAccess.WriteData()
	DataAccess.ReadConnection()

	fmt.Println("file name", DataAccess.FileName)

	//access function using end package name
	DB.ReadSQLDB()
	//File.ReadCSVFile()

	//package aliases used if end package name same (else give error) or if too big path
	dbconfig.ReadServerConfig()
	fileconfig.ReadConfig1()

	//third party packages
	fmt.Println(stringutil.Reverse("PSL"))

	//access func of different file
	//have to give file name of the function while running the main file
	//ex: go run HelloWorld.go Support.go
	fmt.Println(GetSupportData())
}
