package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Bayu",
		"address": "Indonesia",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
