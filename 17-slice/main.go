package main

import "fmt"

func main() {
	//! 1. Slice from explicit Array
	arr := [6]string{"This", "is", "a", "Go", "Interview", "question"}
	fmt.Println(arr)

	s:=arr[1:4] // slice , length= 3, capacity=5
	fmt.Println(s)

	//! 2. slice from slice
	s1 := s[1:3]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//! 3. Slice from literal
	slice := []int{1,2,3}
	fmt.Println("Slice Literal: ",slice)
	fmt.Println("Length: ",len(slice))
	fmt.Println("Capacity: ",cap(slice))

	// ! 4. Slice from make function 
	makeSlice := make([]int, 3) // [0,0,0], length = 3, capacity = 3
	fmt.Println("Make Slice: ",makeSlice)
	fmt.Println("Length: ",len(makeSlice))
	fmt.Println("Capacity: ",cap(makeSlice))

	makeSlice[0] = 5  // [5,0,0], length = 3, capacity = 3
	fmt.Println("Make Slice: ",makeSlice)
	fmt.Println("Length: ",len(makeSlice))
	fmt.Println("Capacity: ",cap(makeSlice))

	makeSlice2:= make([]int, 3, 5) // [0,0,0], len=3, capacity=5
	makeSlice2[0]=5 // [5,0,0], len=3, capacity=5
	fmt.Println("Make Slice 2: ",makeSlice2)
	fmt.Println("Length: ",len(makeSlice2))
	fmt.Println("Capacity: ",cap(makeSlice2))
	// makeSlice2[2]=40 //  runtime error, index out of range as len is 3 for this slice

	// ! 5. Slice from scratch 
	var emptySlice []int // [], empty slice or nil slice
	fmt.Println("Empty Slice: ",emptySlice)
	fmt.Println("Length: ",len(emptySlice))
	fmt.Println("Capacity: ",cap(emptySlice))
	// assign empty slice 
	emptySlice = append(emptySlice, 1);
	fmt.Println("Empty Slice: ",emptySlice)
	fmt.Println("Length: ",len(emptySlice))
	fmt.Println("Capacity: ",cap(emptySlice))
	
	emptySlice = append(emptySlice, 2,3,4,5);
	fmt.Println("Empty Slice: ",emptySlice)
	fmt.Println("Length: ",len(emptySlice))
	fmt.Println("Capacity: ",cap(emptySlice))

	fmt.Println("=========================")
	var x []int  // [], len = 0, cap = 0
	x = append(x, 1) // [1], len = 1, cap = 1 
	fmt.Println(x, len(x), cap(x))
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))
	y:= x 
	fmt.Println("Y: ",y, len(y), cap(y))

	x = append(x, 4)
	fmt.Println("X: ",x, len(x), cap(x))
	fmt.Println("Y: ",y, len(y), cap(y))
	// Slice underlying array rule => 1024 -> 100% increase, 1024 < 25% increase
	y = append(y, 5)
	fmt.Println("X: ",x, len(x), cap(x))
	fmt.Println("Y: ",y, len(y), cap(y))

	x[0] = 10
	fmt.Println("X: ",x, len(x), cap(x))
	fmt.Println("Y: ",y, len(y), cap(y))
	fmt.Println("=========================")

	newSlice := []int{1,2,3,4,5} // [1,2,3,4,5], len = 5, cap = 5
 	fmt.Println(newSlice, len(newSlice), cap(newSlice)) 
	newSlice = append(newSlice, 6) // [1,2,3,4,5,6], len=6,cap=10
	fmt.Println(newSlice, len(newSlice), cap(newSlice)) 
	newSlice = append(newSlice, 7) // [1,2,3,4,5,6,7], len=7,cap=10
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	a:= newSlice[4:] // [5,6,7], len=3, cap=6
	fmt.Println("a", a, len(a), cap(a))

	newY:= changeSlice(a) 
	fmt.Println("NewSlice: ", newSlice, len(newSlice),cap(newSlice)) // [1,2,3,4,10,6,7], len=7,cap=10
	fmt.Println("NewY: ", newY, len(newY),cap(newY)) // [10,6,7,11], len=4,cap=6
	fmt.Println("NewSlice: ", newSlice[0:8], len(newSlice),cap(newSlice)) // [1,2,3,4,10,6,7,11], len=7,cap=10


	print(5,6,7,8,9)
}

func changeSlice(s []int) []int{
 s[0]=10 // [10,6,7], len=3, cap=6
 s = append(s, 11) // [10,6,7,11], len=4,cap=6
 return s
}

// Variadic function
func print(numbers ...int){
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