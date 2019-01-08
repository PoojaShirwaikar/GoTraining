package main

import "fmt"

type Person struct {
	Name string
	Id   int
}

//reference type receiver
func (per *Person) GetPersonDetails(code string) {
	fmt.Println(per.Id, per.Name)
	per.Id = 7777
}

func main() {
	//per1 is of value type
	per1 := Person{"Sachine", 10}
	//per1.Name
	//per1.id

	fmt.Println("before: ", per1)
	per1.GetPersonDetails("test")
	fmt.Println("after: ", per1) //modified id

	//per1 is of value ref
	per2 := new(Person)
	per2.Id = 181
	fmt.Println("before: ", per2)
	per2.GetPersonDetails("test")
	fmt.Println("after: ", per2) //modified id

	//if receiver is of type reference, it will automatically take care of passing by ref/value at runtime

}
