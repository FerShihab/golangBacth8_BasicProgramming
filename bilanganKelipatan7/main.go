package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&num)

	if num%7 == 0 {
		fmt.Printf("%d adalah bilangan kelipatan 7.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan kelipatan 7.\n", num)
	}
}
