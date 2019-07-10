package main

import "fmt"

func main() {
	/* Pointers are aliases */
	original := 0
	copyOriginal1 := &original
	copyOriginal2 := &original
	copyOriginal3 := &original

	*copyOriginal1++
	fmt.Println(original)

	*copyOriginal2++
	fmt.Println(original)

	*copyOriginal3++
	fmt.Println(original)

	/* Creating a variable that is an address already using new(T) */
	number := new(int)   //this is an address of int
	*number = 100        //passing a value
	fmt.Println(*number) //printing the value
	fmt.Println(number)  //printing the address
}
