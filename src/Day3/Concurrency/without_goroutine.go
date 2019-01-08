package main

import "fmt"
import "time"

func main() {

	//synchronous execution.. func2 waits till func1 executes
	func1()
	func2()
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func1", i)
	}
}

func func2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func2", i)
	}
}
