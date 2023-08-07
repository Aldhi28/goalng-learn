package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 1

	// // single condition looping 
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 1; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	/////////////////////////////////////

	// i := 2
	// for{
	// 	fmt.Println("loop ke %d \n",i)

	// 	if i == 0 {
	// 		break
	// 	}
	// 	i--
	// 	time.Sleep(1 * time.Second)
	// }

	///////////////////////////////////////////////
	
		var count int
		fmt.Scanln(&count)
		for i := count; i >= 1; i-- {
			fmt.Println(i, " I will become a go developer")
		}

		/////////////////////////////////////////////////

// 	for n:=0; n<=10; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		fmt.Println("bilangan ganjil %d \n", n)
// 	}
 }