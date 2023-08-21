/*
Pada tugas kali ini kamu diminta untuk mengumpulkan nama-nama yang memiliki jumlah karakter genap dari daftar nama yang diberikan. 


Function description:

Lengkapilah function evenNames dalam editor. Function evenNames memiliki parameter berupa slice string dan harus mengembalikan nilai berupa slice string juga.


For example:

Input	Result
Herman budi jenny kevin aris
Herman budi aris
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
 
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
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    x := scanner.Text()
    slice := strings.Split(x, " ")
    names := evenNames(slice)
    for _, name := range names {
        fmt.Print(name + " ")
    }
    
}
