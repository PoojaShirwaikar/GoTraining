package main

import "fmt"

func main() {
	var array1 [3]int //array of integer of length 3 is created
	//allocate space for 3 values
	//arrays accessed using index - start from 0

	//array1[0] = 10   //by default 0 value
	array1[1] = 20
	array1[2] = 30

	//array1[3] = 10   //invalid - array index out of bound

	fmt.Println(array1)

	//declare and initialise
	var array2 = [2]string{"java", "c#"}

	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	//initialize array within loop
	var array3 [2]int
	for j := 0; j < 2; j++ {
		array3[j] = j + 1
	}
	fmt.Println(array3)

	//iterate array using range loop
	//range is like an iterator
	for index, value := range array3 {
		fmt.Println("index is :", index)
		fmt.Println("value is :", value)
	}
}
