package main

import "fmt"

//map : works with key value pair
func main() {
	var map1 = map[string]int{"key1": 10, "key2": 20, "key3": 30}
	fmt.Println(map1)

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	//var map2 map[int]string //nil map
	//map2[1]="java"   //index out of bound

	var map2 = make(map[int]string)

	map2[100] = "java"
	map2[200] = "c#"
	fmt.Println(map2)

	//add to map
	map2[300] = "golang"
	fmt.Println(map2)

	//update in map
	map2[300] = "angular"
	fmt.Println(map2)

	//delete from map
	delete(map2, 300)
	//delete(name of map, key to delete)
	fmt.Println(map2)

	//can use sort package to sort key/values of map

	//check existance of key within map
	//value, exist := map2[200]

	if value, exist := map2[200]; exist {
		fmt.Println("key exist: ", value)
	} else {
		fmt.Println("key doesnt exist")
	}
}
