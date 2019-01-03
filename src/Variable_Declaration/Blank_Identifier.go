package main

import "fmt"

//blank identifier for package imported
import _ "github.com/golang/example/stringutil"

func main() {

	var name = "pooja"
	var id = 101
	var status = false

	//blank identifier
	_ = status
	//status declared but not used

	//can be used for locally defined variables and globally imported packages

	fmt.Printf(name, id)
}
