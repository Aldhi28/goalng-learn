/*
Pada tugas kali ini kamu diminta untuk membuat kontrak antara struct Segitiga dan interface BangunDatar yang sudah disediakan.

Untuk melakukan kontrak dengan sebuah interface, sebuah tipe data harus memilki (mengimplementasikan) setiap method yang dideklarasi interface tersebut.

Example :

package main

import "fmt"

type Hewan interface {
    bergerak()
}

type Ayam struct {
    code int
    bulu bool
}

func (a *Ayam) bergerak()  {
    fmt.Println("ayam bergerak")
}


For example:

Input	Result
4		  16
8
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
type BangunDatar interface {
    Luas() int
}
 
type Segitiga struct {
    alas   int
    tinggi int
}
 
func (l *Segitiga) Luas() int {
    return l.alas * l.tinggi / 2
}
 
func getLuas(bangunDatar BangunDatar) {
    fmt.Printf("Luas bangun datar = %d \n", bangunDatar.Luas())
}
 
func main() {
    var segitiga1 Segitiga
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    segitiga1.alas, _ = strconv.Atoi(scanner.Text())
    scanner.Scan()
    segitiga1.tinggi, _ = strconv.Atoi(scanner.Text())
    getLuas(&segitiga1)
 
}