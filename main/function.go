package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("hello %s %s \n", firstName, lastName)
}

func increment(angka int) int{
	return 1 + angka
}

func main() {
	sayHello("aldhi","Rizkiyansah")
	angka := increment(1)
	fmt.Println(angka)

}
