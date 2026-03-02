package mathlib

import "fmt"

var money =1000 // global variable can be unused. compiler skip this

var Poysa =0.00

func Subs(x int, y int) {
	subs := x - y

	fmt.Println("Subs",subs)

	// fmt.Println("from mathlib",money)

}