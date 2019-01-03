package main

import "fmt"

func main() {
	//slice: array without specifying slice
	//internally refers to some array only

	var slice1 = []int{1, 2, 23, 4}
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	for _, value := range slice1 {
		//	fmt.Println("index is :", index)
		fmt.Println("value is :", value)
	}

}
