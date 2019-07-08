package main

import "fmt"

/*
Maps in golang are same with object with javascript
*/

func main() {
	// syntax: map[datatype of the key]datatype_of_the_value{}
	var person = map[string]string{}
	person["Name"] = "Vincent"
	person["Age"] = "22"
	person["Gender"] = "Male"
	fmt.Println(person)

	var person1 = map[string]string{
		"Name":   "Vincent",
		"Age":    "22",
		"Gender": "Male",
	}
	fmt.Println(person1)

	/*
		To allow a map to accept a value that has different data types we can use interface
	*/
	var anything = map[string]interface{}{
		"Name":   "Vincent",
		"Age":    22,
		"Gender": "Male",
		"Single": true,
	}
	fmt.Println(anything)

	/*
		If you want to convert an interface,
		you can convert it because you cannot assign interfaces
		to int even though its number nor string
	*/
	var a string = anything["Name"].(string)
	fmt.Println(a)
	var b int = anything["Age"].(int)
	fmt.Println(b)
}
