package main

import (
	"fmt"

	"example.com/mathlib" // custom package
)

// global scope

var a=10
var b=20


func add(x int, y int){
	// local scope
	z:=x+y
	fmt.Println(z)

	sum:= x+a
	fmt.Println(sum)

	fmt.Println()
}

func main() {
	// local scope
	fmt.Println(a,b);

	p:=30
	q:=40

	add(p,q)

	add(a,b)

	add(p,b)

	age:=20

	if age>=20{
		// local scope
		p:=1

		fmt.Println("I've",p,"girlfriend") // get p=1 value due to new scope
	}

	// package scope 
	// go mod init <module-name>
	// go mod init example.com
	mathlib.Subs(p,q)

	fmt.Println("from main package",mathlib.Poysa)

}