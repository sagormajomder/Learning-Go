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

}

/*
Slice contains 3 things 
-> pointer
-> length
-> capacity : starting from pointer location to the last element of the array
*/