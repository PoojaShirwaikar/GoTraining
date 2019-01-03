package main

import "fmt"

func main() {
	var slice1 []int //nil slice
	//dont allocate memory

	//slice1[0] = 100  //give index out of bound error
	//can add element using append
	slice1 = append(slice1, 10)
	fmt.Println(slice1)

	//instead of creating nil slice use make slice

	var slice2 = make([]int, 2) //length and capacity would  be 2
	//allocate memory, size specified
	slice2[0] = 100
	slice2[0] = 200
	fmt.Println(slice2)
}
