package main

import "fmt"

/*
Recursion is calling a function itself,
Recursion can cause StackOverflow (or infinite loop)
To avoid Stackoverflow you must set a condition to make it stop
*/

func main() {
	recursion(0)
}

func recursion(i int) int {
	if i >= 5 {
		fmt.Println("Recursion is done!")
		return i
	}
	i++ // increasing the value to meet the condition and to avoid infinite loop
	fmt.Println(i)
	recursion(i) // recursion calling itself inside of its own body
	return i
}
