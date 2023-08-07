package main

import (
	"flag"
	"fmt"
)

func main() {
	// var nama string
	// fmt.Printf("Nama Lengkap: ")
	// fmt.Scanf("%s", &nama)
	// fmt.Println(nama)


// input alamt sebagai string
	// fmt.Println("Alamat : ") // input alamat
	// scanner := bufio.NewScanner(os.Stdin) // syntax scanner input alamat di atas
	// scanner.Scan() 
	// homeAddress := scanner.Text() 
	// fmt.Println("Alamat Anda \n", homeAddress) // hasil input alamat

// input per array
	// argsWithProg := os.Args //input yang di masukan go run userinput.go aldhi rizki yansyah
	// fmt.Println(argsWithProg[1]) // aldhi
	// fmt.Println(argsWithProg[2]) // rizki
	// fmt.Println(argsWithProg[3]) // yansyahs

	fullNameArg := flag.String("Name", "", "User Name") // -Nama=aldhi
	homeAddressArg := flag.String("adress", "", "User Address") // -adress=lebak bulus
	isgoldMemberArg := flag.Bool("gold", false, "if User Gold Member") // -gold=true

	flag.Parse()
	// cara menggunakan flag
	fmt.Println("Full Name",*fullNameArg) //  hasil input di atas masuk kesini
	fmt.Println("Home Address",*homeAddressArg) // hasil input di atas masuk kesini
	fmt.Println("Gold Member",*isgoldMemberArg) // hasil input di atas masuk kesini
} 