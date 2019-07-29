package main

import "fmt"

func insertName(name string) {
	/* We created a defer with a recover inside so
	that we can still run this function, try to run
	this with defer and without defer to see the
	effect */
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("this function was recovered")
		}
	}()
	if name == "" {
		panic("no name provided")
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
