package main

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
}
