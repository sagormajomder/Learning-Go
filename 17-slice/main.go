package main

import "fmt"

func main() {
	//! 1. Slice from explicit Array
	arr := [6]string{"This", "is", "a", "Go", "Interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // slice , length= 3, capacity=5
	fmt.Println(s)

	//! 2. slice from slice
	s1 := s[1:3]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//! 3. Slice from literal
	// in this case, len and cap always same
	slice := []int{1, 2, 3}
	fmt.Println("Slice Literal: ", slice)
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity: ", cap(slice))

	// ! 4. Slice from make function
	makeSlice := make([]int, 3) // [0,0,0], length = 3, capacity = 3
	fmt.Println("Make Slice: ", makeSlice)
	fmt.Println("Length: ", len(makeSlice))
	fmt.Println("Capacity: ", cap(makeSlice))

	makeSlice[0] = 5 // [5,0,0], length = 3, capacity = 3
	fmt.Println("Make Slice: ", makeSlice)
	fmt.Println("Length: ", len(makeSlice))
	fmt.Println("Capacity: ", cap(makeSlice))

	makeSlice2 := make([]int, 3, 5) // [0,0,0], len=3, capacity=5
	makeSlice2[0] = 5               // [5,0,0], len=3, capacity=5
	fmt.Println("Make Slice 2: ", makeSlice2)
	fmt.Println("Length: ", len(makeSlice2))
	fmt.Println("Capacity: ", cap(makeSlice2))
	// makeSlice2[20]=40 //  runtime error, index out of range as len is 3 for this slice

	// ! 5. Slice from scratch (empty slice or nil slice)
	var emptySlice []int // [], empty slice or nil slice, It doesn't have any underlying array.
	fmt.Println("Empty Slice: ", emptySlice)
	fmt.Println("Length: ", len(emptySlice))
	fmt.Println("Capacity: ", cap(emptySlice))
	// assign empty slice
	emptySlice = append(emptySlice, 1)
	fmt.Println("Empty Slice: ", emptySlice)
	fmt.Println("Length: ", len(emptySlice))
	fmt.Println("Capacity: ", cap(emptySlice))

	emptySlice = append(emptySlice, 2, 3, 4, 5)
	fmt.Println("Empty Slice: ", emptySlice)
	fmt.Println("Length: ", len(emptySlice))
	fmt.Println("Capacity: ", cap(emptySlice))

	fmt.Println("=========================")
	var x []int      // [], len = 0, cap = 0
	x = append(x, 1) 
	fmt.Println(x, len(x), cap(x)) // [1], len = 1, cap = 1
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x)) // [1 2], len = 2, cap=1*2 = 2
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x)) // [1 2 3], len = 3, cap = 2*2 = 4
	y := x // they will ref same underlying array
	fmt.Println("Y: ", y, len(y), cap(y)) // [1 2 3], len = 3, cap = 4

	/*
		numbers:= make([]int, 5, 7)
	numbers= append(numbers, 782, 889, 698)
	শুরুতে
	👉 length = 5
	👉 capacity = 7

	মানে ৫টা element আছে, মোট জায়গা আছে ৭টার।

	append করার সময় আমরা ৩টা নতুন element যোগ করছি,
	কিন্তু ফাঁকা জায়গা ছিল মাত্র ২টা ❌

	তাই Go নতুন একটা array তৈরি করে এবং সাধারণত ছোট slice হলে
	👉 capacity দ্বিগুণ করে।

	পুরনো capacity = 7
	নতুন capacity = 7 × 2 = 14 ✅

	এই কারণেই final output এ cap(numbers) = 14 হয়।

	👉 capacity এর মধ্যে append হলে পুরনো array
	👉 capacity ছাড়িয়ে append হলে নতুন array (capacity বাড়ে).

	Slice underlying array rule =>
	👉 array has 1024 element or less than -> 100% increase or double of prev cap, 
	👉 more than 1024 element -> 25% increase of prev cap
	*/

	x = append(x, 4)
	fmt.Println("X: ", x, len(x), cap(x)) // [1 2 3 4], len = 4, cap = 4
	// `y` showing diff result cause it length is not updated yet. It is still 3.
	fmt.Println("Y: ", y, len(y), cap(y)) // [1 2 3], len = 3, cap = 4


	y = append(y, 5) // overwrite `x` 4th element cause its len still 3
	fmt.Println("X: ", x, len(x), cap(x))  // [1 2 3 5], len = 4, cap = 4
	fmt.Println("Y: ", y, len(y), cap(y)) // [1 2 3 5], len = 4, cap = 4

	x[0] = 10
	fmt.Println("X: ", x, len(x), cap(x)) // [10 2 3 5], len = 4, cap = 4
	fmt.Println("Y: ", y, len(y), cap(y)) // [10 2 3 5], len = 4, cap = 4
	fmt.Println("=========================")

	newSlice := []int{1, 2, 3, 4, 5} // [1,2,3,4,5], len = 5, cap = 5
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 6) // [1,2,3,4,5,6], len=6,cap=5*2=10
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 7) // [1,2,3,4,5,6,7], len=7,cap=10
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	a := newSlice[4:] // [5,6,7], len=3, cap=6
	fmt.Println("a", a, len(a), cap(a))

	newY := changeSlice(a)
	fmt.Println("NewSlice: ", newSlice, len(newSlice), cap(newSlice))      // [1,2,3,4,10,6,7], len=7,cap=10
	fmt.Println("NewY: ", newY, len(newY), cap(newY))                      // [10,6,7,11], len=4,cap=6
	fmt.Println("NewSlice: ", newSlice[0:8], len(newSlice), cap(newSlice)) // [1,2,3,4,10,6,7,11], len=7,cap=10, not give error cause it has value

	print(5, 6, 7, 8, 9)
}

func changeSlice(s []int) []int {
	s[0] = 10         // [10,6,7], len=3, cap=6
	s = append(s, 11) // [10,6,7,11], len=4,cap=6
	return s
}

// Variadic function (a slice)
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

/*
Slice contains 3 things
-> pointer
-> length
-> capacity : starting from pointer location to the last element of the array
*/

//! Simulation

// ************* Compile Phase **************
// **Code segment**
/*
	main = func () {....}
	changeSlice = func(s []int){...}
	print=func(numbers ...int){...}
*/

// **Data Segment**
/*
 */
