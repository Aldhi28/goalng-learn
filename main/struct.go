package main

import "fmt"

type Student struct {
	Name string 
	Alamat string 
	Nim int 
}

func (student Student) newstudent() {
	fmt.Printf("hallo aku %s adalah murid baru \n", student.Name)


}

func main() {
	var student Student
	student.Name = "aldhi"
	student.Nim = 101010
	student.Alamat = "lebak bulus"
	fmt.Println(student)
	fmt.Printf("nama saya %s memiliki nim %d \n ", student.Name, student.Nim)

	aldhi := Student{Name: "aldhi"}
	aldhi.newstudent()
}