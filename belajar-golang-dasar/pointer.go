package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1
	
	// address2.City = "Bandung"

	// fmt.Println(address1) // {Subang Jawa Barat Indonesia}
	// fmt.Println(address2) // {Bandung Jawa Barat Indonesia}

	
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // {Bandung Jawa Barat Indonesia}
}

// kesimpulan 
// - Ketika kita membuat variabel baru dengan tipe data struct, secara otomatis akan menyalin semua nilai dari variabel yang kita buat.
// - Ketika kita mengubah nilai dari salah satu variabel, nilai dari variabel lain tidak akan berubah.
// - Hal ini dikarenakan variabel tersebut memiliki alamat memori yang berbeda.

// Pointer
// - Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada.
// - Pointer digunakan agar program tidak boros memory dan meningkatkan performa.
// - tanda * digunakan untuk membuat pointer.
// - tanda & digunakan untuk mengambil alamat memory.

