package main

import "fmt"
import "time"

func server1(ch1 chan string) {
	time.Sleep(9)
	fmt.Println("server1")
	ch1 <- "response from server1"

}

func server2(ch2 chan string) {
	time.Sleep(6)
	fmt.Println("server2")
	ch2 <- "response from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	//for {
	//when u want to listen for more than one channel
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
		break

	case msg2 := <-ch2:
		fmt.Println(msg2)
		break
	}
	//}
}
