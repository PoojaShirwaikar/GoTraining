package main

import "fmt"

func main() {
	ProcessData1()
	fmt.Println("return back to main")
}

func ProcessData1() {
	defer CloseFile()
	fmt.Println("Processing Started...")
	//throw exception
	//exited abruptly
	panic("error in processdata")
	fmt.Println("Processing ended...")
}

func CloseFile() {
	fmt.Println("do CleanUp....")
}
