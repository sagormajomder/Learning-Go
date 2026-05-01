package main

import "fmt"

// Custom Type (read only)
// we can write first letter of custom types' name and property in both uppercase or lowercase.
// It is just personal style
// But use uppercase as most developer prefer that and it also make it able to export to other packages
type User struct {
	Name string // member variable or property
	age  int
}

type employee struct{
	name string
	salary float64
}

func main() {

	var user1 User
	user1 = User{ // Instance
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

	emp1:=employee{ // Instance
		name: "Sagor",
		salary: 50000,
	}

	fmt.Println(emp1)

}

//! Simulation 

// ************* Compile Phase **************
// **Code segment**
/*
		User = type User struct{...}
		employee = type employee struct{...}
		main = func(){...}
*/

// **Data Segment**
/*
	
*/