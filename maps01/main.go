package main

import "fmt"

func main() {
	// looping throught map with integer key we can also use string as a key ;)
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	for i, _ := range m {
		fmt.Println(m[i])
	}

	// we can put variable as key as long as the type of variable is same with type
	mm := map[string]string{}
	str := "Name"
	mm[str] = "vincent"
	for i, _ := range mm {
		fmt.Println("The key is: ", i, " The value of this key is", mm[i])
	}
}
