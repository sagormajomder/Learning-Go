package main

import "fmt"

//* Defer -> last in First Out (LIFO)
//not immediately call, store them somewhere. defer runs after the return values are set, but before the function actually exits back to the caller.

/*
The Execution Order
1. Evaluate & Assign: The return statement evaluates the expression and assigns the result to the return parameters.

2. Defer: The deferred functions are executed (in Last-In-First-Out order).

3. Exit: Control is handed back to the calling function.
*/

/*
Q: where defer functions stored? what name it is?
Ans: defer functions store as linked Lists but behaves like stack. This linked lists store either heap segment or stack segment. their memory addresses are referenced by `defer list pointer` which is a memory cell of stack segment.
Q: If is name is Stack, then what the diff between memory stack segment and this stack?
Q: If it is store in stack segment, then how it is store into stack segment?
ans: defer list pointer
*/

/*

named return values
-------------
1. all codes execute
2. defer function store kora hobe
3. return -> all defer functions execute korbe
4. return korbe named variables gular value

just return types
------------------
1. all codes execute
2. defer function store kora hobe
3. return values are evaluated at this time (store the return value)
4. all defer functions execute korbe

*/

func a() {
	i := 0

	fmt.Println("First", i) // 0

	defer fmt.Println("Second", i) // fmt.Println("Second",0)
	i++
	fmt.Println("third", i)       // 1
	defer fmt.Println("forth", i) // fmt.Println("forth", 1)
	return
}

// * here `result` is named return value
func sum(a, b int) (result int) {
	result = a + b
	return
}

func calc() int {
	result := 0
	fmt.Println("First", result) // 0

	show := func() {
		result += 10
		fmt.Println("Defer", result) //  fmt.Println("Defer",15)
	}

	defer show()

	result = 5
	fmt.Println("Second", result) // 5
	return result
}

func calculate() (result int) {

	fmt.Println("First", result) // 0

	show := func() {
		result += 10
		fmt.Println("Defer", result) //  fmt.Println("Defer",15)
	}

	defer show()

	result = 5
	fmt.Println("Second", result)
	return
}
func calculate2() (result int) {

	fmt.Println("First", result) // 0

	show := func() {
		result += 10
		fmt.Println("Defer", result) //  fmt.Println("Defer",15)
	}

	defer show()

	result = 5
	p := func(a int) {
		fmt.Println("Ami", a) // 5
	}
	defer p(result)

	defer fmt.Println(result) //5

	fmt.Println("Second", result) // 5

	defer fmt.Println(5)

	return
}

func main() {
	a()
	fmt.Println("=================")
	res := sum(3, 4)
	fmt.Println(res)
	fmt.Println("=================")
	totalCalc := calc()
	fmt.Println("final Calc", totalCalc)
	fmt.Println("=================")
	total := calculate()
	fmt.Println("final Total", total)
	fmt.Println("=================")
	calc2 := calculate2()
	fmt.Println("final calc2", calc2)
}

//! Simulation

// ************* Compile Phase **************
// **Code segment**
/*
	a = func(){...}
	sum = func(a,b int){...}
	calc = func(){...}
	calcAnonymous = func(){...} // bind with calc
	calculate = func(){...}
	calculateAnonymous = func(){...}  // bind with calculate
	calculate2 = func(){...}
	calculate2Anonymous1 = func(){...}  // bind with calculate2
	calculate2Anonymous2 = func(a int){...}  // bind with calculate2
	main = func () {....}
*/

// **Data Segment**
/*
 */
