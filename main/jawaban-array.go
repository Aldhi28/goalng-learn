/*
Doni ingin mengetahui ada berapa angka genap yang terdapat di sebuah array, hanya saja doni butuh bantuan untuk mengetahuinya karena elemen di dalam array tersebut sangat banyak. Tulislah code untuk menerima 2 inputan dari Doni,

- Inputan pertama adalah kapasitas dari array

- Inputan kedua adalah elemen dari array tersebut.

Lalu cetak angka-angka genap tersebut ke dalam console.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
}
