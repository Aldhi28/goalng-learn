package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {


	/*
	Doni ingin mengetahui ada berapa angka genap yang terdapat di sebuah array, hanya saja doni butuh bantuan untuk mengetahuinya karena elemen di dalam array tersebut sangat banyak. Tulislah code untuk menerima 2 inputan dari Doni, 

	- Inputan pertama adalah kapasitas dari array 

	- Inputan kedua adalah elemen dari array tersebut. 

	Lalu cetak angka-angka genap tersebut ke dalam console.
	*/
     // jawaban dari pertanyaan di atas
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    capacity, _ := strconv.Atoi(scanner.Text())
    arr := make([]int, capacity)

    scanner.Scan()
    arrText := scanner.Text()
    arrText2 := strings.Split(arrText, " ")

    for i, v := range arrText2 {
        x, _ := strconv.Atoi(string(v))
        arr[i] = x
    }

    for _, v := range arr {
        if v%2 == 0 && v != 0 {
            fmt.Println(v)
        }
    }


// for ranges array = untuk perulangan element array memberikan informasi index ke berapa selain nilai nilai element nya

	// direction := [4]string {"utara", "selatan", "timur", "barat"}
	// for i, direction := range direction {
	// fmt.Printf("element %d : %s \n ", i, direction)
	// }


// adeskord atau sepesial carakter untuk mengabaikan index
	// direction := [4]string {"utara", "selatan", "timur", "barat"}
	// for _, direction := range direction {
	// fmt.Printf("element %s \n ", direction)
	// }


// array multi dimensional
// len = untuk mendapat informasi atau kampasitas array nya

	// var num = [4][3] string {
	// 	[3]string {"1", "2", "3"},
	// 	[3]string {"4", "5", "6"},
	// 	[3]string {"7", "8", "9"},
		
	// }
	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("baris ke %d , isinya %v \n", i+1, num[i])
	// }


	// penulisan array secara vertical

	// tipeProduk := [...]string{
	// 	"Barang",
	// 	"Jasa",
	// 	"Jasa dua",
	// 	"Jasa tiga",
	// 	"Jasa empat",
	// }

	// tipeProduk[4] = "Jasa lima" // untuk mengubah isi array yang sudah di tentukan
	// fmt.Println(tipeProduk)

	// deklarasi variable
	// var tipeproduct [2]string = [2]string{"Barang", "Jasa"}
	// fmt.Println(tipeproduct)

	// var tipeproduct = [2]string("Barang","Jasa")
	// tipeproduct := [2]string{"Barang", "Jasa"}

	/////////////////////////////////////
	// tipeproduct[0] = "Barang"
	// tipeproduct[1] = "Jasa"

	// fmt.Println(tipeproduct[0])
	// fmt.Println(tipeproduct[1])
	// fmt.Println(tipeproduct)

}