package main

import "fmt"

// custom value
type User struct{
	Name string
	Age int 
	salary float64
}

// pass by value
func print(numbers [3]int){
	fmt.Println("Copy of Arr", numbers)
}

// pass by reference
func printPointer(numbers *[3]int){
	fmt.Println("Actual Array",*numbers)
}

func main() {
	// pointer or address of memory
	x := 20

	p := &x // ampersand(&) = address of

	fmt.Println("Value: ",x) // value
	fmt.Println("Address: ",p) // address of x 
	// * -> value at address
	fmt.Println("Value at address: ",*p) // value of p address

	x = 30
fmt.Println("Value: ",x) // value

*p = 40; // change value at address p

fmt.Println("Value at address: ",*p) // value of p address
fmt.Println("Value: ",x) // value


arr:= [3]int{1,2,3}
print(arr)

printPointer(&arr)

sagor:= User{
	Name: "Sagor",
	Age: 29,
	salary: 0.0,
}

pSagor := &sagor

fmt.Println(sagor)
fmt.Println(pSagor)
fmt.Println(pSagor.Name)


}