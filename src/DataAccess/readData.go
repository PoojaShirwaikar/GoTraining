//reusable package

//package name same as container folder name
package DataAccess

import "fmt"

var FileName string = "file1.txt"
var fileLocation string = "c:\\temp" //not exported outside package (1st letter should be capitalized)

func init() {

	fmt.Println("init from read data called")
}

func ReadData() {
	fmt.Println("readData")
}

func WriteData() {
	var config = readConfig()
	fmt.Println("writeData", config)
}
