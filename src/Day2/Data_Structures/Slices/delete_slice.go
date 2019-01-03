package main

import "fmt"

func main() {
	myslice := []string{"monday", "tuesday"}
	myotherslice := []string{"wednesday", "thursday", "friday"}
	myslice = append(myslice, myotherslice...)
	fmt.Println(myslice)

	//delete element at 2nf index
	//slicing + append
	myslice = append(myslice[:2], myslice[3:]...)
	fmt.Println(myslice)
}
