package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	// var resultInt int = result.(int) // error

	fmt.Println(resultString)
	// fmt.Println(resultInt)
}