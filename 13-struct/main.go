package main

import "fmt"

// Custom Type
type User struct {
	Name string // member variable or property
	age  int
}

func main() {

	var user1 User
	user1 = User{
		Name: "Sagor",
		age:  29,
	}
	fmt.Println(user1)

	user2:= User{ // Instance
		Name: "Mehedi",
		age: 30,
	}
	fmt.Println("Name: ", user2.Name)
	fmt.Println("age: ", user2.age)

}