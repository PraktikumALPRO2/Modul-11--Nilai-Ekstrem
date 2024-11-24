package main

import (
	"fmt"
)

func main() {
	var berat_145 [1000]float64

	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci antara 1 dan 1000.")
		return
	}

	fmt.Println("Masukkan berat anak kelinci (spasi):")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat_145[i])
	}

	minBerat := berat_145[0]
	maxBerat := berat_145[0]

	for i := 1; i < n; i++ {
		if berat_145[i] < minBerat {
			minBerat = berat_145[i]
		}
		if berat_145[i] > maxBerat {
			maxBerat = berat_145[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
