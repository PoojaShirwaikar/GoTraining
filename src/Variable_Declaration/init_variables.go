package main

import "fmt"

//global variables
var name string = "Pooja"
var id int
var status bool
var salary float32

//name = "Pooja"  //cannot be done

//local variables declared within function
func main() {
	//standard variable declaration

	/*
		var name string = "Pooja"
		var id int = 123
		var status bool = true
		var salary float32 = 123.67
	*/

	name = "Pooja"
	id = 123
	status = true
	salary = 123.67

	fmt.Println(name, id, status, salary)
}
