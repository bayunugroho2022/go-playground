package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // {Bandung Jawa Barat Indonesia}
}

// kesimpulan
// asterisk operator digunakan untuk membuat pointer. 
// Ketika kita ingin mengubah nilai dari pointer, kita perlu menambahkan tanda * sebelum nama variabelnya.
