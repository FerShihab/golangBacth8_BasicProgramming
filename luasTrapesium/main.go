package main

import (
	"fmt"
)

func calculateTrapezoidArea(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

func main() {
	var base1, base2, height float64
	fmt.Print("Masukkan panjang alas atas: ")
	fmt.Scan(&base1)
	fmt.Print("Masukkan panjang alas bawah: ")
	fmt.Scan(&base2)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&height)

	area := calculateTrapezoidArea(base1, base2, height)
	fmt.Printf("Luas trapesium: %.1f\n", area)
}
