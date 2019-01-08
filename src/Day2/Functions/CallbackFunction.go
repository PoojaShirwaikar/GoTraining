package main

import "fmt"

func IterateNumbers(slice1 []int, callback func(no int)) {

	for _, value := range slice1 {
		callback(value)
		//instead of having global variables to access the values.. a function defination is passed as parameter
	}
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	IterateNumbers(slice1, func(no int) {
		fmt.Println(no)
	})
}
