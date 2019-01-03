package main

import "fmt"

func main() {
	//slice: array without specifying slice
	//internally refers to some array only

	//slice is of type structure
	//1 pointer to internal array
	//2 length of internal array
	//3 capacity

	var slice1 = []int{1, 2, 23, 4}
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	for _, value := range slice1 {
		//	fmt.Println("index is :", index)
		fmt.Println("value is :", value)
	}

	//appending new element at end of slice
	slice1 = append(slice1, 6, 7, 8, 9)
	fmt.Println(slice1)
}
