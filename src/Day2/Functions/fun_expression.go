package main

import "fmt"

func main() {
	//function Expression
	greet := func() {
		fmt.Print("function restricted within main")
	}

	greet()
}
