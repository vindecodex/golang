package main

import "fmt"

//Note that we can attach methods to any type we created
// Think that this struct is our class
type People struct {
	// Something like a properties of our class People
	Talk string
}

// This is our methods of our People class
func (p People) talk() {
	fmt.Println(p.Talk)
}

func main() {
	// We instansciate class People and also giving a value to our property Talk which is string
	var goku = People{"Spirit Ball!!!"}
	goku.talk()
}
