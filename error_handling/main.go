package main

import (
	"errors"
	"fmt"
)

//Creating our custom error message using error object
func students(name ...string) ([]string, error) {
	if len(name) <= 0 {
		// creating new error message using errors.New("Message")
		return nil, errors.New("empty: no names entered")
	}
	return name, nil
}

func main() {
	result, err := students(
		"Vincent",
		"John",
		"Jacob",
		"Trisha",
	)

	// panic will halt the program
	// panic is not ideal it is recommended to use object error instead
	// but i just added it here to tackle error handling because it is also part of it
	if err != nil {
		defer fmt.Println("defer here, this must be top of the panic, we can also run a function here :) ")
		panic(err.Error()) //Printing the error that we created but the program will be halted
	}

	// after using a panic you must have a defer, because defer will be
	// the last to run after all the function done running

	fmt.Println(result)

	// to activate our error, remove the arguments that we passed on students function
}
