package main

import (
	"fmt"
)
 
func main() {
    var kumpulanSekatBamboo []int
    var banyakBamboo int
    var cuts int
 
   
    fmt.Scanln(&banyakBamboo)
    for i := 0; i < banyakBamboo; i++ {
        var SekatBamboo int
        fmt.Scanln(&SekatBamboo)
        kumpulanSekatBamboo = append(kumpulanSekatBamboo, SekatBamboo)
    }

    fmt.Scanln(&cuts)
 
    for i := 1; i <= cuts; i++ {
        fmt.Println("Potongan ke-", i)
        kumpulanSekatBamboo = decreaseBamboos(kumpulanSekatBamboo)
        printBamboos(kumpulanSekatBamboo)
    }
 
}
 
func decreaseBamboos(bamboos []int) []int {
    for i := 0; i < len(bamboos); i++ {
        if bamboos[i] == 0 {
            continue
        }
        bamboos[i] -= 1
    }
    return bamboos
}
 
func printBamboos(bamboos []int) {
    for i := 0; i < len(bamboos); i++ {
        printBamboo(bamboos[i])
    }
}
 
func printBamboo(bamboo int) {
    fmt.Println(bamboo)
}