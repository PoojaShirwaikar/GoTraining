package main

import "fmt"

func main() {
	defer foo() //called just before end of main
	//mostly used for clean up activities

	bar()

	value := fetchData()
	_ = value
}

func foo() {
	fmt.Println("foo called")
}

func bar() {
	fmt.Println("bar called")
}

func fetchData() string {
	defer cleanup() // executed before returning from function

	status := true

	if status == true {
		return "connected"
	} else {
		return "error"
	}
}

func cleanup() {
	fmt.Println("cleanup called")
}
