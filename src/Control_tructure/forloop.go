package main

import "fmt"

func main() {
	//pattern 1:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//pattern 2:
	//dont have a while loop in go
	//this pattern similar to While loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j := j + 2
	}

	//pattern 3:
	//infinite loop
	//can be used with break stmt
	k := 0
	for {
		fmt.Println(k)
		k := k + 2

		if k > 10 {
			break
		}
	}

	//pattern 4:
	//with break and continue
	var m = 0

	for {
		m++

		if m/2 == 0 {
			continue
		}
		fmt.Println(m)
		if m > 20 {
			break
		}
	}

}
