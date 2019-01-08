package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Name        string
	Id          int
	notExported int
}

func main() {

	var user1 = User1{"Ajay", 100, 10}
	//marshal: convert go structure to json
	bs, _ := json.Marshal(user1)

	fmt.Println(bs)
	//fmt.Printf("%T \n",bs)
	fmt.Printf(string(bs))
}
