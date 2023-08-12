package main

import "fmt"

type Address struct {
	City, Provice string
}

func templateAddress(address *Address) {
	address.Provice = "Jawa tengah"
}

func main() {
	address := Address{
		City: "Semarang",
		Provice: "",
	}

	var addressPointer *Address = &address
	templateAddress(addressPointer)
	fmt.Println(address)
}