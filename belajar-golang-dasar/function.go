package main

import "fmt"


func getAllName() string {
	return "Bayu Nugroho";
}

func mutiplValue() (string, string, string) {
	return "Bayu", "Nugroho", "Indonesia"
}

func nameReturnValue() (firstName string, lastName string) {
	firstName = "Bayu"
	lastName = "Nugroho"

	return
}

func main() {

	// single value
	name := getAllName()
	fmt.Println(name)

	// multiple value
	firstName, lastName, _ := mutiplValue()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// name return value
	firstName, lastName = nameReturnValue()
	fmt.Println(firstName)
	fmt.Println(lastName)
	
}

