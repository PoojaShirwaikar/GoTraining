package main

import "fmt"

func main() {
	//shorthand variable declaration
	name, id, status := "Pooja", 10, true

	fmt.Printf("%T,%T,%T", name, id, status)
}
