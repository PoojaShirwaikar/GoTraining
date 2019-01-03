package main

import "fmt"

func main() {
	//shorthand variable declaration
	name := "Pooja"
	id := 10
	status := true

	//can declare only inside function
	//cant declare globally

	fmt.Printf("%T,%T,%T", name, id, status)
}
