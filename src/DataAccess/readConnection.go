package DataAccess

import "fmt"

//exported outside package
func ReadConnection() {
	fmt.Println("readConnection")
}

//function accesible within package
func readConfig() string {
	return "config func from readConnection"
}
