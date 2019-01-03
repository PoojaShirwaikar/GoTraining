package main

import "fmt"
import "sort"

func main() {
	myslice := []int{1, 2, 3}
	myslice1 := make([]int, 4)

	copy(myslice1, myslice)
	//copy(destination,source)
	fmt.Println(myslice1)

	//sort integers
	sort.Ints(myslice)

	slice1 := myslice[:]
	fmt.Println(slice1)

	//create slice from array
	var array1 = [5]int{1, 2, 3, 4, 5}
	//slice1:=array[1:4]
	slice := array1[:]
	fmt.Println(slice)
}
