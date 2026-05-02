package main

import "fmt"

func main() {
	// Variable
	// var a int = 10
	// var a = 10
	// a := 40.5
	// a := "Hello World"
	a := true
	a = false

	const p = 100

	fmt.Println(a)
	fmt.Println(p)

	var x int = 10
	fmt.Printf("%d \n", x) // decimal = d

	var y float64 = 10.635
	fmt.Printf("%.2f \n", y) // float = f

	r := '💖'
	fmt.Printf("%c \n", r) // character = c

	var isBool bool = true
	fmt.Printf("%v \n", isBool)

	var s string = "My name is Sagor"
	fmt.Printf("%s \n",s)

	//know the type
	fmt.Printf("** type of variable s = %T \n",s)
	fmt.Printf("** type of variable isBool = %T \n",isBool)
}

/**
%d
--------
signed
int8 = 8bits = 1 byte ; 2^8 = 256 = -128 to 127
int16 = 16bits = 2 byte;
int32 = 32bits = 4 byte;
int64 = 64 bits = 8 byte;

unsigned
uint8 = 8bits = 1 byte ; 2^8 = 256 = 0 to 255
uint16 = 16bits = 2 byte;
uint32 = 32bits = 4 byte;
uint64 = 64 bits = 8 byte;

%f
---------------
float32 = 32bits = 4 byte;
float64 = 64 bits = 8 byte;

%v
--------------
bool = 8 bits = 1byte;

byte = alias of uint8 = 8 bit per character = 1 byte
rune = alias for int32 (unicode point) = 32 bits = 4 byte = %c (interview q)

*/
