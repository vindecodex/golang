package main

import "fmt"

func b() []interface{} {
	A := func() string {
		return "AAAA"
	}
	B := func() int {
		return 200
	}
	C := func() string {
		return "CCCC"
	}
	return []interface{}{
		A(),
		B(),
		C(),
	}
}

func main() {
	fmt.Println(b()[0])
	fmt.Println(b()[1])
	fmt.Println(b()[2])
}
