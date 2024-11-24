package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	if n > 1000 {
		fmt.Println("Jumlah anak kelinci tidak boleh lebih dari 1000")
		return
	}

	var berat float64
	minBerat := math.MaxFloat64
	maxBerat := -math.MaxFloat64

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat)

		if berat < minBerat {
			minBerat = berat
		}
		if berat > maxBerat {
			maxBerat = berat
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
