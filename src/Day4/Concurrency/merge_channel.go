package main

import (
	"fmt"
	"time"
)

func Server1(ch chan int) {
	time.Sleep(3)
	//ch <- 1000
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func Server2(ch chan int) {
	time.Sleep(5)
	//ch <- 2000
	for i := 3; i < 6; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Server1(ch1)
	go Server2(ch2)

	merge(ch1, ch2)

}

func merge(ch ...chan int) {
	for value, data := range ch {
		//fmt.Println("values are: ", value, <-data)

		//if looping through in sender.. have to have a looping in merge
		for i := 0; i < 3; i++ {
			fmt.Println("values are: ", value, <-data)
		}
	}
}
