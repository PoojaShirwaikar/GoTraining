package main

import "fmt"

var Counter int

//global function
//can be called from any file in package
//even if local..
//any functions within package can access the function

// func Increment() int {
// 	Counter++
// 	return Counter
// }

//Requirement.....
//Counter should be private variable
//but should be initialized omly once

var nextCounter = func() (func() int, func() int) {
	var counter int //counter initialised only once
	var increment = func() int {
		counter++
		return counter
	}

	var decrement = func() int {
		counter--
		return counter
	}

	return increment, decrement
}

func main() {
	cntr, cntr1 := nextCounter() //called once...

	//fmt.Println(nextCounter()()) //nextCounter() called  each time... counter initialised each time

	fmt.Println(cntr()) //the function defination in nextCounter called
	fmt.Println(cntr())
	fmt.Println(cntr1())

	// fmt.Println(Increment())
	// fmt.Println(Increment())
	// fmt.Println(Increment())
	// fmt.Println(Increment())
	// fmt.Println(Increment())
}
