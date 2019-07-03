package main

import "fmt"

func main() {
	a := 100 //variable a holding 100
	b := &a  //variable b pointing to a or holding an address of a
	fmt.Printf("a = %v b = %v \n", a, b)
	//outputs: a = 100 b = 0xc00008c000%

	*b = 200 //changing the value of a to 200
	fmt.Printf("a: %v", a)
	// outputs: a = 200
}
