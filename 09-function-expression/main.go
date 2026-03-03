package main

import "fmt"

func main() {

	a:= 10 // variable expression

	// function expression
	add:= func (a, b int){ 
		fmt.Println(a+b)
	}

	add(a,10)
}


func init() {
	fmt.Println("I will be calling first")
}