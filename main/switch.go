package main

import "fmt"

func main() {
	tipeProduk := "Barang" // memanggil case yang ada di bawah

	switch tipeProduk {
	case "Barang":
		fmt.Println("Memiliki Informasi Stock")
	case "Jasa":
		fmt.Println("Tidak Memiliki stock awal")
	default:
		fmt.Println("Tipe Produk Tidak Ada")
	}

	var nilai = 80

	switch {
	case nilai > 80:
		fmt.Println("A")
	case nilai > 70:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}