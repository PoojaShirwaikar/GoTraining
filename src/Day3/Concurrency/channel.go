package main

import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int)  //unbuffered channel



	wg.Add(2)  //2 go routines
	go func1() //go routine 1
	go func2() //go routine 2
	wg.Wait()  //wait till go routines return done
}


//put something in channel
func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func1", i)
		fmt.Println("value put channel ", ch1<-i)
	}
	wg.Done() //required else main will not know that goroutine is executed
}

//retrieve from channel
func func2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func2", i)
		fmt.Println("value retrived from channel ", <-ch1)

	}
	wg.Done()
}
