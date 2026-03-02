package main

import "fmt"

var a = 10
var b = 15

func main() {
	age := 29

	if age >= 18 {
		a := 47 // same name but different variable as we redeclare means variable shadowing
		b = 20 // as reassign so same variable
		fmt.Println("A: ",a)
		fmt.Println("B: ",b)
	}

	fmt.Println("A: ",a)
	fmt.Println("B: ",b)
}