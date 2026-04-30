package main
import "fmt"

const a=10 // store in code segment as it is read only value

var p=100

func call(){
	// add() function also in code segment in the compile phase.
	// call() function just reference it in its stack frame in the execution phase.
	// As nested function, add() is bind with call(). So only call() function can access add() function. 
	// In compile phase, It just a anonymous function which is bind with call() function. 
	// The anonymous function store into add variable in execution phase, not compile phase.
    add:=func (x int, y int){
        z:=x+y
        fmt.Println(z)
    }
    
    add(5,6)
    add(p,a)
}

func main() {
  call()
  fmt.Println(a)
}

func init(){
    fmt.Println("I will be executed first")
}

/*

2 phases for the Go Program running:
	1. compilation phase
	2. execution phase

	--> go run main.go => compile--> create a binary file named main--> then automatically execute the  binary file
	--> go run main.go => compile => main => ./main

	--> go build main.go => compile--> create a binary file named main --> then we can execute the binary file by running ./main
	--> go build main.go => compile => main

	--> Go is a compiled language, so it compiles the code before executing it.
	-->In the compilation phase, all the constants and functions are allocated in the code segment(read-only). And all the global variables are allocated in the data segment.
	--> In the execution phase, the stack is used for function execution and local variables.

*/

/*
-> program execution step by step
	i) compilation phase
		- syntax checking (main function check, rules and regulation check etc)
		- produce executable/binary file

		
	ii) execution phase

	code seg - load the compiled file code segment data to RAMs code segment

	data seg - check for global and static variables, if present allocate memory in data segment and initialize them

		- check for init() function, if present execute it by allocating memory in stack segment called stack frame
		***(Stack frame - memory allocated for a function execution in stack segment)
		
		- check for main() function, if present execute it by allocating memory in stack segment called stack frame
		
		- check for other function calls inside main() and execute them by allocating memory in stack segment called stack frame recursively
		- deallocate the memory in stack segment in LIFO manner after function execution is completed

	data lookup process -> local scope(stack segment) -> heap -> global scope(data segment) -> code segment

	*/
	
//! Simulation 

// ************* Compile Phase **************
// **Code segment**
/*
		a=10
		call = func() {....}
		callAnonymous = func () {....}
		main = func () {....}
		init = func () {....}
*/

// **Data Segment**
/*
	p=100
*/

	