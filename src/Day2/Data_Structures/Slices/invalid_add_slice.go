package main

import "fmt"

func main() {
	greetings := make([]string, 3, 5) //allocate nil value for 3 position
	//	greetings[0] = "hi"
	//	greetings[1] = "good morning"
	//	greetings[2] = "good night"

	fmt.Println(greetings)
	//greetings[3] = "hello"  //index out of bound
	//to add elements beyound length use append

	greetings = append(greetings, "hello") //appends to 4th location (index 3)
}
