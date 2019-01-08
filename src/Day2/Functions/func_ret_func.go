package main

import "fmt"

func main() {
	func1 := func11() //func1 will hold definition of greet()
	fmt.Println(func1())
}

//func which will return function
func func11() func() int {
	// return func() int {
	// 	return 100
	// }

	greet := func() int {
		return 100
	}

	return greet
}
