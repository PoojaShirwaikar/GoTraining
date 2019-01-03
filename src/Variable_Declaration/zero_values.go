package main

import "fmt"

//local variables declared within function
func main() {
	//standard variable declaration
	var name string
	var id int
	var status bool
	var salary float32

	var (
		a, b int
		c, d string
	)

	fmt.Println(name, id, status, salary)
}
