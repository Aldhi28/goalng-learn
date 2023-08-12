/*
Bantulah Susi untuk menentukan jadwal sehari-harinya dengan sebuah program sederhana. Akan diberikan inputan jam berupa integer lalu tampilkanlah pada console kegiatan yang akan dilakukan Susi sesuai jam yang diberikan. Kegiatan yang akan dilakukan Susi adalah berikut :

Jika jam 4-5 akan mencetak “Bangun Pagi”

Jika jam 6-7 akan mencetak “Mandi dan Sarapan”.

Jika jam 8-11 akan mencetak “Berangkat Sekolah”.

Jika jam 12 akan mencetak “Pulang Sekolah”.

Jika jam 13-15 akan mencetak “Tidur Siang”.

Jika Jam 16-17 akan mencetak “Waktu Main”

Selain dari itu akan mencetak “Waktu Belajar dan Istirahat”

Bila user menginput melebihi 24 jam maka akan mencetak “Hanya ada 24 jam dalam sehari”


Sample input 1:
9

Sample Input 2:
25

Sample Input 3:
23



Sample output 1:
Berangkat Sekolah

Sample Output 2:
Hanya ada 24 Jam dalam sehari

Sample Output 3:
Waktu Belajar dan Istirahat
*/

package main

import "fmt"

func main() {
var jam int
	fmt.Scanln(&jam)

	if jam > 24 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	} else if jam == 4 || jam == 5 {
		fmt.Println("Bangun Pagi")
	} else if jam == 6 || jam == 7 {
		fmt.Println("Mandi dan Sarapan")
	} else if jam >= 8 && jam <= 11 {
		fmt.Println("Berangkat Sekolah")
	} else if jam == 12 {
		fmt.Println("Pulang Sekolah")
	}
}