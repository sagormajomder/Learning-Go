package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are eligible to be married")
	}else if age <= 10{
		fmt.Println("You are just a teenage, not eligible to be married")
		
	}else {
		fmt.Println("You are not eligible to be married, but you can love someone")
	}

	a:=2

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1 nor 2 or 3")
	}
}