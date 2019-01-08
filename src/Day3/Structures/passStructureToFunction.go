package main

import "fmt"

type Product struct {
	ProdName string
	ProdId   int
}

func main() {
	//create structure by value
	var prd1 = Product{ProdName: "TV", ProdId: 100}
	fmt.Println("before: ", prd1)
	ModifyStruct(&prd1)
	fmt.Println("after: ", prd1)

}

func ModifyStruct(prod *Product) {
	prod.ProdName = "modified tv"
}
