package main

import "fmt"

func main() {
	// expr1 := func() {
	// 	fmt.Println("self executing")
	// }

	// expr1()

	//called when declared
	//executed only once
	//cannot be manually called
	func() {
		fmt.Println("self executing function")
	}()
}
