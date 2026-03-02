package main

import "fmt"

var a = 10

func main() {
	fmt.Println("Hello Init Function")
	fmt.Println(a)
}

func init(){
	fmt.Println(a)
	a = 20
	fmt.Println("I am the first function that is executed first")
}