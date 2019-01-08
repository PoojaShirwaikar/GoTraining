package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go ProcessData(ch1)

	//when using range to iterate through channel.. it doesnt know where to stop so it goes on iteration... get deadlock error
	// require the sender to close() channel
	for value := range ch1 {
		fmt.Println(value)
	}
}

func ProcessData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	// require the sender to close() channel when using range function to iterate
	close(ch1)
}
