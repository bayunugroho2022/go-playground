package main

import "fmt"


func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	fmt.Println("Terjadi error:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}

