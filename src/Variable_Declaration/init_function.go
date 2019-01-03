package main

import "fmt"
import "DataAccess"
import "reflect"
import "strconv"

var a, b, c int

//can have more than one init functions
//it gets called only once
//called before running main funtion
//used in case of one time initialization
//can be used within package also
//cant be called explicitly
func init() {
	a = 100
	b = 200
	fmt.Println(a, b)
}

func init() {

	fmt.Println("second init called")
}

func main() {
	fmt.Println("in main")

	//if readData.go has another package import than that init called first
	//init functions called from behind
	//even if package not used (blank identifier) the init works
	//init functions from a file executed together
	//files from package chosen randomly
	DataAccess.WriteData()

	//reflect package has function to find type of variable
	fmt.Println(reflect.TypeOf(a))

	//strconv package has function for string to int conversion
	result, error := strconv.Atoi("1234")
	_ = error
	fmt.Println(result)

	//strconv package has function for int to string conversion
	result1 := strconv.Itoa(1234)

	fmt.Println(result1)

}
