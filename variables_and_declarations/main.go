package main

import "fmt"

func des(dez ...interface{}) {}

func main() {
	//syntax : var variable data_type
	var a int

	//declaring more variables with similar data types
	var b, c int //this is called parallel declaration

	//declaring more variables with different data types
	var (
		d int
		e string
		f bool
		g float64
	)

	//assigning a value on a variable short way
	var h, i, j int = 1, 2, 3

	//Another shortway that has different types
	k, l := 10, "ways"

	// dont mind this line, i just add it so that no errors will happen
	// because golang will get an error if declaring a variable without using it
	des(a, b, c, d, e, f, g, h, i, j, k, l)

	/* Declaring nil array, slice and maps */

	// var arr [5]int
	// if arr == nil {
	// 	fmt.Println("arr is nil")
	// }
	/* this array will not be a nil because array accepts length, by adding
	length indexes will be fill out by 0 or blank strings depends on type we use */

	var slice []int
	if slice == nil {
		fmt.Println("slice is nil")
	}

	var mp map[string]int
	if mp == nil {
		fmt.Println("map is nil")
	}

	mp = map[string]int{
		"1": 1,
		"2": 2,
	}

	fmt.Println(mp)

	/*
		Use var to set things to zero
		Pointers, Maps, Slices, Functions, Interfaces, channels can be nil
	*/

}
