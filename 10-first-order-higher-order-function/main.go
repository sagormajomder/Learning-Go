package main

import "fmt"

// Higher order function / first class function
// 1. func as parameter
func processOperation(a int, b int, op func (x int, y int)){
	op(a,b)
}
// 2. return func
func call () func (a,b int) {
	return add
}

/*
1. first order function
	- standard or named function
	- anonymous function
	- IIFE
	- function expression
*/

func add(a, b int){
	c:= a+b 
	fmt.Println(c)
}

func main(){
	processOperation(10,20,add)

	sum:= call()

	sum(2,7)
}