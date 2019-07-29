package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

var slice = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	odd := filter(slice, isOdd)
	even := filter(slice, isEven)

	fmt.Println("slice = ", slice)
	fmt.Println("Odd elements of slice are: ", odd)
	fmt.Println("Even elements of slice are: ", even)
}
