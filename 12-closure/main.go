package main

import "fmt"

const a = 10

var p=100

func outer() func() {
	money:=100 //closure value
	age := 30

	fmt.Println("age is", age)

	// In compile phase, It just a anonymous function which is bind with outer() function. 
	// The anonymous function store into inner variable in execution phase, not compile phase
	inner := func() {
		money=money+a+p

		fmt.Println(money)
	}

	return inner //closure value
}

func call(){
	incre1 := outer()
	incre1()
	incre1()

	incre2 := outer()
	incre2()
	incre2()
}

func main() {
call()
}

func init(){
	fmt.Println("====Bank===")
}

//! Simulation 

// ************* Compile Phase **************
// **Code segment**
/*
		a=10
		outer = func () {....}
		outerAnonymous = func() {....}
		call = func() {....}
		main = func () {....}
		init = func () {....}
*/

// **Data Segment**
/*
	p=100
*/
