package main

import "fmt"

func main() {
	/*var i = 10
	if i == 0 {
		fmt.Println("in if, i = 0")
	} else if i > 10 {
		fmt.Println("in else if, i >10")
	} else {
		fmt.Println("in else")
	}*/

	//scope of i within the if-else
	if i := getSalary(); i == 0 {
		fmt.Println("in if, i = 0")
	} else if i > 100 {
		fmt.Println("in else if, i >100")
	} else {
		fmt.Println("in else")
	}

}

func getSalary() int {
	return 1000
}
