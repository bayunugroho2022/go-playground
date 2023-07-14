package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	one := Ups(1)
	fmt.Println(one)

	two := Ups(2)
	fmt.Println(two)

	three := Ups(3)
	fmt.Println(three)
}