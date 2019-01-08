package main

import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup

func main() {
	//executing asynchronously
	wg.Add(2)  //2 go routines
	go func1() //go routine 1
	go func2() //go routine 2
	wg.Wait()  //wait till go routines return done
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func1", i)
	}
	wg.Done() //required else main will not know that goroutine is executed
}

func func2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func2", i)
	}
	wg.Done()
}
