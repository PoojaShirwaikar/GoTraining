package main

import "fmt"

func main() {
	var name string = "john"

	//pattern 1
	switch name {
	case "john":
		fmt.Println("john called")

	case "jenny":
		fmt.Println("jenny called")

	default:
		fmt.Println("default")
	}

	//pattern 2
	switch name {
	case "john":
		fmt.Println("john called")
		fallthrough //fallthrough next case

	case "jenny":
		fmt.Println("jenny called")
		fallthrough

	case "ajay":
		fmt.Println("ajay called")

	default:
		fmt.Println("default")
	}

	//pattern 3
	switch name {
	case "john", "jenny":
		fmt.Println("john,jenny called")
		fallthrough //fallthrough next case

	case "ajay":
		fmt.Println("ajay called")

	default:
		fmt.Println("default")
	}

	//pattern 4
	switch {
	case len(name) == 4:
		fmt.Println("length = 4")

	case name == "ajay":
		fmt.Println("name is ajay ")

	case getsalary() == 1000:
		fmt.Println("SALARY  = 1000")
	default:
		fmt.Println("default")
	}

	//SwitchOnType(7)
	SwitchOnType("pooja")
}

func getsalary() int {
	return 1000
}

//x of type empty interface
func SwitchOnType(x interface{}) {
	switch x.(type) { //this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")

	}
}
