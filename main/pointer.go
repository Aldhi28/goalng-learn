package main

import "fmt"

func main() {
	var myVal = 64
	var myValPointer *int = &myVal
	fmt.Println(myVal)
	fmt.Println(myValPointer)
	fmt.Println(*myValPointer)

	var myVal2 *int
	fmt.Println(myVal2)
}