package main

import "fmt"

type val int

// Using a struct with pointers are something like making a storage of that type

func increase(num *val) {
	*num = *num + 10
}

func main() {
	var a1, a2 val
	fmt.Println(a1, " ", a2)
	//increase accepts a pointer so we give it an address of the variable a1
	increase(&a1)
	fmt.Println(a1, " ", a2)
	increase(&a2)
	fmt.Println(a1, " ", a2)
}
