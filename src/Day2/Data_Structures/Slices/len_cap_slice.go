package main

import "fmt"

func main() {
	var slice2 = make([]int, 2) //length and capacity would  be 2
	//var slice2 = make([]int, 2, 8) //length = 2 and capacity =8
	slice2[0] = 100
	slice2[1] = 200
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))
	slice2 = append(slice2, 300)
	//capacity exhausted so it doubles capacity... internally discard the array and create new array of capacity 4
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 400)
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 500)
	//again capacity exhausted so it doubles capacity...
	fmt.Println(len(slice2), cap(slice2))

	slice3 := append(slice2, 500)
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)

	slice2[4] = 777
	//changed in slice2 and slice3  -- passed by reference
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)

	fmt.Println(len(slice2), cap(slice2))
	slice4 := append(slice2, 600, 700, 800, 900, 1000)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(len(slice4), cap(slice4))
	slice2[4] = 154665
	//doesnt change 4th position in slice4
	//capacity exhausted for slice4 and new array created internally so slice2 n slice4 point to different locations
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
}
