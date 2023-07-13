package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var noKtpBayu NoKTP = "123456789"
	var marriedStatus Married = false

	fmt.Println(noKtpBayu)
	fmt.Println(marriedStatus)
}

