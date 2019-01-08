package main

import "fmt"

func main() {
	ch1 := make(chan int, 3) //buffered channel
	//buffer: how many elements it can have in channel without blocking channel

	ch1 <- 12
	ch1 <- 13
	ch1 <- 14

	fmt.Println("value of channel ", <-ch1)
	fmt.Println("value of channel ", <-ch1)

	//can add more values after some values retrived
	ch1 <- 15

}
