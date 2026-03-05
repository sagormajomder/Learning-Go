package main

import "fmt"

// Custom Type
type User struct {
	Name string // member variable or property
	age  int
}

func printUserDetails(usr User){
		fmt.Println("Name: ", usr.Name)
	fmt.Println("age: ", usr.age)
}
// Receiver function (only work with struct)
func (usr User) printDetails(){
		fmt.Println("Name: ", usr.Name)
	fmt.Println("age: ", usr.age)
}

func (usr User) call(a int){
	fmt.Println("User Name:",usr.Name);
	fmt.Println("User Age:", a)
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
	
	printUserDetails(user2)

	// call receiver function
	user2.printDetails()
	user1.call(user1.age)

}

/*

**Compile Phase**
-------------------

*code segment*
--------------
type User struct {.....}
printUserDetails: func (usr User){....}
printDetails: func(){....} // binding with User Type variable
call: func (a int){...} // binding with User Type variable
main: func (){...}
*/