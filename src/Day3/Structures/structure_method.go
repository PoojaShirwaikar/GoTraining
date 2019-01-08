package main

import "fmt"

type Person struct {
	Name string
	Id   int
}

//method for Person struct
//function with receiver receiver
//can have only one receiver
//member method of Person struct
func (per Person) GetPersonDetails() {
	fmt.Println(per.Id, per.Name)
}

//function
func GetPersonDetails() {
	fmt.Println("function")
}

func main() {
	per1 := Person{"Sachine", 10}
	//per1.Name
	//per1.id

	//method called
	per1.GetPersonDetails()

	//function called
	GetPersonDetails()

	//cannot have method or function overloading
}
