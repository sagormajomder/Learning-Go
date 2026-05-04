package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func add(x,y int){
	fmt.Println(x+y)
}

func printHello(num int) {
	fmt.Println("Hello! Sagor", num)
	add(2,4)
}

func main() {
	x:= 10
	fmt.Println("Hello", x)
	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, " ", p)
 
	time.Sleep(5 * time.Second)
}

//! Simulation

// ************* Compile Phase **************
// **Code segment**
/*
	p=11
	add = func(x,y int){...}
	printHello = func(num int){...}
	main = func() {....}
*/

// **Data Segment**
/*
	a=10
*/
