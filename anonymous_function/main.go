package main

import "fmt"

func main() {
	// my anonymous func
	func() {
		fmt.Println("This is anonymous")
	}()
}
