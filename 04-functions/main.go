package main

import "fmt"

// No returns
func add (a int, b int) {
	sum:=a+b
	fmt.Println(sum)
}

// With returns
func add2 (a int, b int) int {
	return a+b
}

// multiple returns
func divide (a int, b int)(int, int){
	result:=a/b
	remainder:=a%b

	return result, remainder
}

func main() {
	a := 10
	b := 20
	// sum := a + b
	// fmt.Println(sum)

	add(a,b)

	sum:=add2(a,b)
	fmt.Println("ADD2:", sum)

	result, remainder := divide(a,b)

	fmt.Println("Divide Result", result)
	fmt.Println("Remainder", remainder)
	fmt.Println()

	// A Application
	app()
	
}