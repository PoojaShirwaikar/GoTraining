package main

import "fmt"

type User struct {
	Name string
	Id   int
}

type Admin struct {
	//embedd another structure here
	User        //Reusability
	role string //cant be used outside package
}

func (usr User) getUserDetails() {
	fmt.Println(usr.Id, usr.Name)
}

func main() {
	//allow some kind of inheritance using embedded types
	//embedd i=one structure to another
	var user1 = User{"Ajay", 100}

	var admin1 = Admin{user1, "Admin"}
	fmt.Println(admin1.Id)
	fmt.Println(admin1.Name)
	fmt.Println(admin1.role)
}
