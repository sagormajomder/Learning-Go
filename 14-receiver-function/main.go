package main

import "fmt"

// Custom Type
type User struct {
	Name string // member variable or property
	age  int
}

type Employee struct{
	Name string
	Salary float64
}

func printUserDetails(usr User){
		fmt.Println("Name: ", usr.Name)
	fmt.Println("age: ", usr.age)
}
// Receiver function (only work with struct)
// fully same as normal function syntax. just write `struct name and type` before function name. 
func (usr User) printDetails(){
		fmt.Println("Name: ", usr.Name)
	fmt.Println("age: ", usr.age)
}

func (usr User) call(a int){
	fmt.Println("User Name:",usr.Name);
	fmt.Println("User Age:", a)
}

func main() {

	var user1 User // Instance
	user1 = User{ 
		Name: "Sagor",
		age:  29,
	}
	fmt.Println(user1)

	user2:= User{ // Instance
		Name: "Mehedi",
		age: 30,
	}
	
	printUserDetails(user2)

	// invoke receiver function
	user2.printDetails()
	user1.call(user1.age)

	//* only User instances can invoke printDetails and call
	// call(20) // give compile error

	emp1 := Employee{
		Name: "Sagor",
		Salary: 50000.250,
	}
	fmt.Println(emp1);

	// emp1.call(20) // give compile error

}

//! Simulation 

// ************* Compile Phase **************
// **Code segment**
/*
		type User struct {.....}
		type Employee struct {.....}
		printUserDetails: func (usr User){....}
		printDetails: func(){....} // binding with User Type variable
		call: func (a int){...} // binding with User Type variable
		main: func (){...}
*/

// **Data Segment**
/*
	
*/