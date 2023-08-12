package main

import "fmt"

// interface kosong
func Hi(nama string) interface{} {
	return "hi " + nama
	} 
	
type Student interface {
	StudentName() string
}

type Biodata struct {
	Nama string
}

func (biodata Biodata) StudentName()  string {
	return biodata.Nama
}

func getStudent(student Student) {
	fmt.Println("Hello", student.StudentName())
}

func main() {
	var aldhi Biodata
	aldhi.Nama = "rizkiyansah"
	getStudent(aldhi)

	var kosong interface{} = Hi("aldhi")
	fmt.Println(kosong)
}