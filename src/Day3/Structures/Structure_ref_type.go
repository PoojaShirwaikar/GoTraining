package main

import "fmt"

type Item struct {
	ItemName string
	ItemId   int
}

func main() {
	//value type
	//item1:=Item{"Laptop",9876}
	//fmt.Println(item1)

	//creating reference type 1:
	item1 := new(Item)
	fmt.Println(item1) //zero values
	item1.ItemId = 987
	item1.ItemName = "mobile"
	fmt.Println(item1)
	//output: &{laptop 987}

	fmt.Println("before: ", item1)
	ModifyItemStruct(item1) //no need to pass address specifically
	fmt.Println("after: ", item1)

	//creating reference type 2:
	item2 := &Item{"Mouse", 22}
	fmt.Println(item2)
}

func ModifyItemStruct(item *Item) {
	item.ItemName = "Samsung"
}
