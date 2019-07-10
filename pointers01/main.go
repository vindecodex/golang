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
}
