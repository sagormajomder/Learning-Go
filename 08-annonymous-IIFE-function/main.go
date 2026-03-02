package main

import "fmt"

// standard or named function
func add(a, b int) {
	fmt.Println(a+b)
}
func main() {
	add(8,10)

	// Anonymous function 
	// IIFE = immediately invoked function expression
	func (x int, y int){
		c:=x+y
		fmt.Println(c)
	}(2,5)
}

func init(){
	fmt.Println("I wil executed first")
}