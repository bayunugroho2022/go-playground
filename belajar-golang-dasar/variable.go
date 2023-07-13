package main

import "fmt"

func main() {

	// option 1
	var name string = "Bayu"

	fmt.Println(name + " from Indonesia")

	// option 2
	fullName := "Bayu Nugroho"
	
	fmt.Println(fullName + " from Indonesia")

	// option 3
	var (
		firstName = "Bayu"
		lastName = "Nugroho"
	)

	
	fmt.Println(firstName + " " + lastName)
}

