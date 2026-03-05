package main

import "fmt"

func outer() func() {
	count := 0

	inner := func() {
		count += 1

		fmt.Println(count)
	}

	return inner
}

func main() {

	closure:=outer()

	closure()
	closure()

}