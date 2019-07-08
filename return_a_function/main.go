package main

import "fmt"

// function that accepts integer that returns a function that will also return interger
func squareRoot(num int) func() int {
	return func() int {
		return num * num
	}
}

func main() {
	// invoking
	fmt.Println(squareRoot(5)())

	// or
	invoke := squareRoot(2)
	fmt.Println(invoke())

}
