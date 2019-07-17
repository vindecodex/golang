package main

import "fmt"

/* Closure 1 */
func a() {
	c := 100
	b(c)
}

func b(num int) {
	num = num + num
	fmt.Println(num)
}

/* Closure 2 */
func c1() func() func() {
	num := 100
	fmt.Println("Initialize variable: ", num)
	return func() func() {
		num = num + num
		fmt.Println("Creates closure level 1 ", num)
		return func() {
			num = num + num
			fmt.Println("Creates closure level 2", num)
		}
	}
}

func main() {
	a()
	c1()()()
}
