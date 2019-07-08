package main

import "fmt"

func main() {
	/*
		Array are sequence of the same data type and has fixed length, but we can make an
		array that is flexible in length
	*/
	/* Fixed length */
	var a [5]int
	fmt.Println("This is a: ", a)

	/* flexible length */
	var b = []int{}
	fmt.Println("This is b: ", b)

	/* flexible length but has default length */
	var c = make([]int, 5)
	fmt.Println("This is c: ", c)

	/*
		The Multi-Dimensional Array
	*/
	var d = [5][5]int{}
	fmt.Println("This is d: ", d)

	/*
		Display Specific indexes on array
	*/
	var text = "MASTER"
	fmt.Println(text[3:4]) // prints T
	fmt.Println(text[1:6]) // prints ASTER
	fmt.Println(text[:3])  // prints MAS
	fmt.Println(text[3:])  // prints TER
	// [lowest index : highest index] or [start : end]
}
