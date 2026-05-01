package main

import "fmt"

// custom value
type User struct {
	Name   string
	Age    int
	salary float64
}

//* pass by value (copy value)
// it take same amount of memory cell of the original array
// for larger array, it take time and huge memory to copy which is major performance drawback
// it is the most reason, why we use pointer
func print(numbers [3]int) {
	fmt.Println("Copy of Arr", numbers)
}

//* pass by reference
// occupy single cell of RAM to store the first index memory address of the array
func printPointer(numbers *[3]int) {
	fmt.Println("Actual Array", *numbers)
}

func main() {
	// pointer means address of memory(RAM)
	x := 20

	//* & -> address of
	p := &x // ampersand(&)

	fmt.Println("Value: ", x)   // value
	fmt.Println("Address: ", p) // address of x
	//* -> value at address
	fmt.Println("Value at address: ", *p) // value of p address

	x = 30
	fmt.Println("Value: ", x) // value

	*p = 40 // change value at address p

	fmt.Println("Value at address: ", *p) // value of p address
	fmt.Println("Value: ", x)             // value

	arr := [3]int{1, 2, 3}
	print(arr)

	printPointer(&arr) //send address of first array index memory cell

	sagor := User{
		Name:   "Sagor",
		Age:    29,
		salary: 0.0,
	}

	pSagor := &sagor

	fmt.Println(sagor)
	fmt.Println(pSagor)
	fmt.Println(pSagor.Name)

}

//! Simulation

// ************* Compile Phase **************
// **Code segment**
/*
	User = type User Struct {....}
	print = func(){....}
	printPointer = func(){......}
	main = func(){....}
*/

// **Data Segment**
/*
	p=100
*/
