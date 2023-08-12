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