package main

import "fmt"

//Structure:
//can hold members of different data types
//EmpName, EmpNo,code,salary
//structure:=  memeber variable + methods

type Employee struct {
	firstname string
	lastname  string
	id        int
}

func main() {

	//pattern 1
	emp1 := Employee{firstname: "Pooja", lastname: "Shirwaiakar", id: 1}
	fmt.Println(emp1)
	fmt.Println(emp1.firstname)
	fmt.Println(emp1.lastname)
	fmt.Println(emp1.id)

	//pattern 2
	//Structures are zero valued
	//set zero values based on data type
	var emp2 Employee
	fmt.Println(emp2.firstname)
	fmt.Println(emp2.lastname)
	fmt.Println(emp2.id)

	//pattern 3
	emp3 := Employee{}
	emp3.id = 2
	emp3.firstname = "p"
	emp3.lastname = "s"
	fmt.Println(emp3)

	//pattern 4
	//without specifying keys
	//specify values in order declared
	emp4 := Employee{"Pooja", "Shirwaiakar", 11}
	fmt.Println(emp4)

}
