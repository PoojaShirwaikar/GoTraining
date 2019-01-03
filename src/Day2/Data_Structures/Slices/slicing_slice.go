package main

import "fmt"

func main() {
	greeting := []string{
		"java",
		"c#",
		"golang",
		"c++",
		"node.js",
		"angular",
	}

	//fmt.Println(greeting[2:4])
	greeting1 := greeting[2:4]
	//include lower limit and (higher limit-1)  in this case doesnt include 4th index
	fmt.Println(greeting1)

	greeting[2] = "changed something....."
	fmt.Println(greeting)
	fmt.Println(greeting1)

	//slicing pattern
	fmt.Println(greeting[:2])
	//if not specified start from 0th index

	fmt.Println(greeting[3:])
	//if not specified considers last index

}
