package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
fungsi len mendapatkan kampasitas element dari slice
fungsi cap mendapatkan informas kampasitas maksimum dari sebuah slice
*/

func evenNames(slice []string) []string {
    var names []string
    for _, v := range slice {
        if len(v)%2 == 0 {
            names = append(names, v)
        }
    }
    return names
}
func main() {
/*
	Pada tugas kali ini kamu diminta untuk mengumpulkan nama-nama yang memiliki jumlah karakter genap dari daftar nama yang diberikan. 

	Function description:

	Lengkapilah function evenNames dalam editor. Function evenNames memiliki parameter berupa slice string dan harus mengembalikan nilai berupa slice string juga.


	For example:

	Input	Result
	Herman budi jenny kevin aris
	Herman budi aris
*/

scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
x := scanner.Text()
slice := strings.Split(x, " ")
names := evenNames(slice)
for _, name := range names {
	fmt.Print(name + " ")
}










	// direction := [] string {"utara", "barat", "timur", "selatan"}
	// var verticalderiction = direction[0:2]

	// fmt.Printf("isi slice %v, jumlah element %d, kampasitas slice : %d, in located %p \n", direction, len(direction), cap(direction),direction)
	// fmt.Printf("isi slice %v, jumlah element %d, kampasitas slice : %d, in located %p \n", verticalderiction, len(verticalderiction), cap(verticalderiction),verticalderiction)
	
	// newVerticalDirection := append(verticalderiction, "Barat laut","Timur laut","Barat daya","tenggara")
	// fmt.Printf("isi slice %v, jumlah element %d, kampasitas slice : %d, in located %p \n", newVerticalDirection, len(newVerticalDirection), cap(newVerticalDirection),newVerticalDirection)

	// fmt.Printf("%p\n", &direction[0])
	// fmt.Printf("%p\n", &verticalderiction[0])
	// fmt.Printf("%p\n", &newVerticalDirection[0])

	// dst := make([]string,3)
	// src := []string{"semangka","apple","jeruk"}
	// copy(dst, src)

	// fmt.Printf("isi src : %v, alamat : %p\n", src, src)
	// fmt.Printf("isi dst : %v, alamat : %p\n", dst, dst)

	///////////////////////////////////////////////////////////////////////
	// direction := [] string {"utara", "barat", "timur", "selatan"}
	// var verticalderiction = direction[2:]

	// fmt.Printf("%v, data type %T\n", direction, direction)
	// fmt.Printf("%v, data type %T\n", verticalderiction, verticalderiction)

	// fmt.Printf("%s, located in %p\n", direction[2], &direction[2])
	// fmt.Printf("%s, located in %p\n", verticalderiction[0], &verticalderiction[0])



	// direction := [] string {"utara", "selatan", "timur", "barat"}
	// fmt.Println(direction)

	// //menggunakan make 
	// s := make([]int,3)
	// fmt.Println("angka",s)



	// direction := [4]string{"utara", "selatan", "timur", "barat"}
	// var verticalderiction = direction[0:2]

	// fmt.Printf("%v , data type %T\n", direction, direction)
	// fmt.Printf("%v , data type %T\n", verticalderiction, verticalderiction)
}