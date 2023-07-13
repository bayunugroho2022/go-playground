package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	
	// Variadic function
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	// Variadic function with slice
	slice := []int{10, 20, 30, 40, 50}
	total = sumAll(slice...)
	fmt.Println(total)

}

