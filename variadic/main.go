package main

import "fmt"

// variadic parameters accepts paramaters 0 or indifinite numbers of parameters
func names(name ...string) []string {
	// variadic parameters becomes a slice so we return type of slice
	return name
}

func main() {
	// we can input indifinite numbers of arguments as long as its type is string
	result := names(
		"John Doe",
		"Jane Doe",
		"John Smith",
		"Jane Smith",
	)
	fmt.Println(result)
}
