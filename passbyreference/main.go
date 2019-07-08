package main

import "fmt"

type Number int

func passByValue(a Number) {
	a = a + 1
}

func passByReference(a *Number) {
	*a = *a + 1
}

func main() {
	var num Number = 0

	passByValue(num)
	fmt.Println(num) // outputs 0

	passByReference(&num)
	fmt.Println(num) // outputs 1
}
