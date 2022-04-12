package main

import "fmt"

func do(i interface{}) {
	switch variable := i.(type) {
	case int:
		fmt.Println(variable)
	case string:
		fmt.Println(variable)
	case func():
		variable()
	default:
		fmt.Println("Default")
	}
}

func main() {

	do(22)
	do("stst")
	do(func() { fmt.Println("aaa") })
}
