/*
Pada tugas ini kamu diminta untuk mencetak kalimat "I will become a go developer" pada console sebanyak inputan yang akan diberikan. Agar tantangan ini menjadi lebih menarik tambahkanlah suatu angka mundur di depan setiap kalimatnya.

Kriteria :

  > Decrement

  > Siapkan variabel untuk menerima inputan dari console



For example:

Input	Result
2
2  I will become a go developer
1  I will become a go developer
*/

package main

import "fmt"

func main() {
    var count int
    fmt.Scanln(&count)
    for i := count; i >= 1; i-- {
        fmt.Println(i, " I will become a go developer")
    }
}
