package main

import "fmt"

func printWelcomeMessage(){
fmt.Println("Welcome to the application")
}

func getUserName() string{
	var name string 
	fmt.Println("Please enter your name")
	fmt.Scanln(&name)
	fmt.Println("User Name", name)
	return name
}

func printGoodbyeMessage(){
	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye")
}

func getTwoNumber() (int, int){
	var num1 int
	var num2 int 
	fmt.Println("Give first Num --")
	fmt.Scanln(&num1)
	fmt.Println("Give second Num - ")
	fmt.Scanln(&num2)
	return num1, num2
}

func displayResult(name string, sum int){
	fmt.Println("Hello ,", name)
	fmt.Println("Summation is ", sum)
}

func app() {
	// welcome message
	printWelcomeMessage()
	// get user name as input
	userName:=getUserName()

	// get two number 
	num1, num2:=getTwoNumber()

	sum:= add2(num1, num2)

	displayResult(userName,sum)

	// print a goodbye message
	printGoodbyeMessage()
}