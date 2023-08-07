package main

import "fmt"

func main () {
// number := 2
// number2 := 2
// if number%2 == 0 && number2%2 == 1 {
// 	fmt.Println("Nilai Genap dan Ganjil")
// } else if number%2 == 0 {
// 	fmt.Println("Nilai Genap")
// } else {
// 	fmt.Println("Nilai Ganjil")
// }

// temporari variable

// if waktu := 3; waktu >= 6 && waktu <= 10 {
// 	fmt.Println("Menu Sarapan")
// }else if waktu > 10 && waktu <= 14 {
// 	fmt.Println("Makan Siang")
// }else {
// 	fmt.Println("Makan Malam")
// }

// var jam int
// fmt.Scanln(&jam)
// var SampleInput = 5
// var SampleInput1 = 9;
// var SampleInput2 = 25;
// var SampleInput3 = 12

// if SampleInput >= 4 && SampleInput <= 5 {
// 	fmt.Println("Bangun Pagi")
// }else if  SampleInput1 >= 8 && SampleInput1 <= 11 {
// 	fmt.Println("Berangkat Sekolah")
// }else if  SampleInput3 == 12 {
// 	fmt.Println("Berangkat Sekolah")
// }else if SampleInput2 >= 25{
// 	fmt.Println("Hanya ada 24 jam dalam sehari")
// }else {
// 	fmt.Println("Waktu belajar dan istirahat")
// }
//////////////////////////////////////////////////


// var jam int
// 	fmt.Scanln(&jam)

// 	if jam > 24 {
// 		fmt.Println("Hanya ada 24 jam dalam sehari")
// 	} else if jam == 4 || jam == 5 {
// 		fmt.Println("Bangun Pagi")
// 	} else if jam == 6 || jam == 7 {
// 		fmt.Println("Mandi dan Sarapan")
// 	} else if jam >= 8 && jam <= 11 {
// 		fmt.Println("Berangkat Sekolah")
// 	} else if jam == 12 {
// 		fmt.Println("Pulang Sekolah")
// 	}

// ///////////////////////////////////////////

// func printActivity(jam int) {
// 	if jam >= 4 && jam <= 5 {
// 		fmt.Println("Bangun Pagi")
// 	} else if jam >= 6 && jam <= 7 {
// 		fmt.Println("Mandi dan Sarapan")
// 	} else if jam >= 8 && jam <= 11 {
// 		fmt.Println("Berangkat Sekolah")
// 	} else if jam == 12 {
// 		fmt.Println("Pulang Sekolah")
// 	} else if jam >= 13 && jam <= 15 {
// 		fmt.Println("Tidur Siang")
// 	} else if jam >= 16 && jam <= 17 {
// 		fmt.Println("Waktu Main")
// 	} else {
// 		fmt.Println("Waktu Belajar dan Istirahat")
// 	}
// }

	var jam int
	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&jam)


	if jam < 0 || jam > 23 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	} else {
			if jam >= 4 && jam <= 5 {
				fmt.Println("Bangun Pagi")
			} else if jam >= 6 && jam <= 7 {
				fmt.Println("Mandi dan Sarapan")
			} else if jam >= 8 && jam <= 11 {
				fmt.Println("Berangkat Sekolah")
			} else if jam == 12 {
				fmt.Println("Pulang Sekolah")
			} else if jam >= 13 && jam <= 15 {
				fmt.Println("Tidur Siang")
			} else if jam >= 16 && jam <= 17 {
				fmt.Println("Waktu Main")
			} else {
				fmt.Println("Waktu Belajar dan Istirahat")
			}
		}

}

