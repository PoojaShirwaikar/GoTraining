package main

import "fmt"

type User1 struct {
	Name string
	Id   int
}

type Admin1 struct {
	//embedd another structure here
	User1        //Reusability
	role  string //cant be used outside package
}

func (usr User1) getUserDetails() {
	fmt.Println(usr.Id, usr.Name)
}

//overriding getUserDetails of parent class
func (admin Admin1) getUserDetails() {
	fmt.Println("from admin ", admin.Id)
}

func main() {
	//allow some kind of inheritance using embedded types
	//embedd i=one structure to another
	var user1 = User1{"Ajay", 100}
	var admin1 = Admin1{user1, "Admin"}

	fmt.Println(admin1.Id)
	//fmt.Println(admin1.Name)
	//fmt.Println(admin1.role)

	admin1.getUserDetails() //calls the overriden method
}
