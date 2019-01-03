package main

import "fmt"

func main() {
	//type inference
	var name = "pooja"
	var id = 101
	var status = false

	//if you want variables to have default value, use standard declaration

	fmt.Printf("%T \n", name) //type of variable
	fmt.Printf(name, id, status)
}
