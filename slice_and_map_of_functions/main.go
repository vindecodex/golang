package main

import "fmt"

func main() {
	var m = make(map[string]func(param string) string)
	m["dispName"] = func(name string) string {
		return name
	}
	fmt.Println(m["dispName"]("vindecode"))

	var arr []func(param string) string
	arr = append(arr, func(name string) string {
		return name
	})
	fmt.Println(arr[0]("vindecode"))
}
