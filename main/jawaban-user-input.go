/*
Tulislah code yang akan menerima inputan berupa string lalu tampilkanlah di console. Inputan yang akan diberikan berupa :

Nama

Alamat

Email


Kriteria :

Siapkan variabel untuk menerima inputan dari console


Sample input 1:
Tian
Jakarta
Tian@gmail.com

Sample output 1:
Halo, saya Tian. Saya tinggal di Jakarta. Alamat email saya adalah Tian@gmail.com



Sample input 2:
Felix
Bandung
Felix5261@yahoo.co.id

Sample output 2:
Halo, saya Felix. Saya tinggal di Bandung. Alamat email saya adalah Felix5261@yahoo.co.id
*/



package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    nama := scanner.Text()
    scanner.Scan()
    alamat := scanner.Text()
    scanner.Scan()
    email := scanner.Text()

    fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s", nama, alamat, email)

}
