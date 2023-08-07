package main

import "fmt"

func main(){
	bulan := map[int]string{
		1:"jan",
		2: "frb",
		3: "mar",
	}
	bulan[4] = "apr" // cara menambahkan value ke variable
	// fmt.Println(bulan)

	var sales = make(map[string]int)
	sales["jan"] = 1 // akan mengecek varible ini
	sales["feb"] = 2
	sales["mar"] = 3
	fmt.Println(sales)

	for k,v := range sales {
		fmt.Printf("sales %s sejumlah %d \n", k,v)
	}

	var value, isExist = sales["jan"] // variable ini akan mengambil nilai yang di atas "jan"
	if isExist {
		fmt.Println(value)
	}else{
		fmt.Println("data tidak di temukan")
	}
}