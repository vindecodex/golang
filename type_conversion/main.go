package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var integer int

	/* Integer --> String */
	//method 1
	integer = 100
	str = fmt.Sprintf("%d", integer)
	fmt.Println(str)
	// method 2
	// Itoa = Integer to ASCII
	str = strconv.Itoa(integer)
	fmt.Println(str)

	/* String --> Integer */
	integer, _ = strconv.Atoi(str)
	fmt.Println(integer)

}
