package main

import "fmt"

func main() {
	//pointer: stores memory address of variable

	var i int = 100

	//passed by value -- copy of variable is sent
	// fmt.Println("value of i before change: ", i) //100
	// change(i)
	// fmt.Println("value of i after change: ", i) //100

	//passed by reference : pass memeory address of a variable
	fmt.Println(&i) //memory address

	var iPointer *int //variable hold pointer to any int variable
	iPointer = &i
	//fmt.Println(iPointer)
	fmt.Println("value of i before change: ", i)
	change(iPointer)
	fmt.Println("value of i after change: ", i)
}

func change(data *int) { //* added so that function accepts pointer
	//we need to fetch data from above memory address
	*data = 900 //dereferencing pointer
}
