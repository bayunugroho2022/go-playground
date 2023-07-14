package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	
	// option 1
	var bayu Customer
	bayu.Name = "Bayu"
	bayu.Address = "Indonesia"
	bayu.Age = 23

	fmt.Println(bayu)

	// option 2
	bayu2 := Customer{
		Name:    "Bayu",
		Address: "Indonesia",
		Age:     23,
	}

	fmt.Println(bayu2)
}
