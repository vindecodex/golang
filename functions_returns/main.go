package main

import "fmt"

// function that returns two values
func a() (int, int) {
	n1 := 10
	n2 := n1 + n1
	return n1, n2
}

// function that returns two values which is name and age
func b() (name string, age int) {
	name = "vincent"
	age = 10
	// no need to specify what we are going to return here becuase we already specify at the top
	return
}

func main() {
	fmt.Println(a())
	fmt.Println(b())
}
