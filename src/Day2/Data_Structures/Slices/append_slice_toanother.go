package main

import "fmt"

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	myotherslice := []int{6, 7, 8, 9}

	myslice = append(myslice, myotherslice...) //veriadic parrameter
	fmt.Println(myslice)
}
