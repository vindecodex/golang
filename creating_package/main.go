package main

import "golang/creating_package/mypackage"

func main() {
	//In Go if you want to make a variable Public you'll just need to Capitalize the first letter of your functions,types,variables etc...
	mypackage.Multiply(10, 10)
}
