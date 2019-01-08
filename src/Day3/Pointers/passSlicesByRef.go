package main

import "fmt"

func main() {

	//how slice are passed by reference

	slice1 := []int{1, 2, 3, 4}
	fmt.Println("before: ", slice1)
	//slice and map by default passed by reference
	changeSlice(slice1)
	fmt.Println("after: ", slice1)

	//arrays are passed by value
	//pass as slice or as pointwer
	array1 := [4]int{1, 2, 3, 4}
	fmt.Println("before: ", array1)
	changearray(&array1)
	fmt.Println("after: ", array1)
}

func changeSlice(slice1 []int) {
	slice1[0] = 78946
}

func changearray(array1 *[4]int) {
	array1[0] = 154 //dont require derefrencing
}
