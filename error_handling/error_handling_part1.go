package main

import (
	"errors"
	"fmt"
)

func insertName(name string) {
	/* We created a defer with a recover inside so
	that we can still run this function, try to run
	this with defer and without defer to see the
	effect */
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("this function was recovered")
			fmt.Println(x) // x is the arguments passed inside panic()
		}
	}()
	if name == "" {
		err := errors.New("no name provided")
		panic(err.Error())
	} else {
		fmt.Println(name)
	}
}

func main() {
	name1 := ""
	name2 := "Vincent"
	insertName(name1)
	insertName(name2)
}
