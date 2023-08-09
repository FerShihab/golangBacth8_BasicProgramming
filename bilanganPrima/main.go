package main

import (
	"fmt"
)

func isBilPrima(n int) bool {
	// bilangan kurang dari atau sama dengan 1 bukan bilangan prima
	if n <= 1 {
		return false
	}

	// sistem looping untuk mengecek bilangan bila kondisi diatas tidak terpenuhi
	// kondisi awal berada di int 2
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&num)

	if isBilPrima(num) {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	}
}
