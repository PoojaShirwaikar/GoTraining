package main

import (
	"fmt"
	"reflect"
)

func main() {
	func1("pooja", "goa")
	id := func3("pooja")
	fmt.Println(id)

	id1, _ := func4("pooja")
	fmt.Println(id1)

	amount, status := getbonus("pooja")
	fmt.Println(amount, status)

	variadicFunc("pooja", "goa")
	variadicFunc("pooja", "goa", 10, 20)

	//pass slice to variadic parameter
	slice1 := []int{1, 2, 3}
	variadicFunc("pooja", "goa", slice1...)

	variadicExample("pooja", 123, true)

	//pass array to variadic parameter
	//cant use array directly.. send as slice
	array1 := [2]int{1, 2}
	variadicFunc("pooja", "goa", array1[:]...)

}

//pattern 1
func func1(name string, address string) {
	fmt.Println(name, address)
}

//pattern 2
//all parameters of string type
func func2(name, address string) {
	fmt.Println(name, address)
}

//pattern 3
func func3(name string) int {
	return 100
}

//pattern 4
func func4(name string) (int, string) {
	return 100, "pooja"
}

//pattern 5
//return named parameter
func getbonus(name string) (amount int, status bool) {
	amount = 1000
	status = true
	//return amount, status //sequence of returning should be same as declared
	return //specifying name with return type so no need to specify again
}

//pattern 6
//variadic parameter function  -- can accept 0 or more values
//can pass 0 or more parameters of type integer
//variadic parameter should be always at end
func variadicFunc(name string, address string, values ...int) {
	//fmt.Println(values)
	fmt.Println(name, address)
	for _, value := range values {
		fmt.Println(value)
	}
}

func variadicExample(i ...interface{}) {
	for _, v := range i {
		fmt.Println(reflect.ValueOf(v).Kind()) //getting data type using reflect
	}
}
