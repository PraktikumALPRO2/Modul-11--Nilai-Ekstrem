package main

import (
	"fmt"
)

func main() {
	var n int
	var weights = make([]float64, 1000) // Array dengan kapasitas 1000

	// Input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci : ")
	fmt.Scan(&n)

	// Input berat masing-masing anak kelinci
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Mencari berat terkecil dan terbesar
	var min, max float64 = weights[0], weights[0]
	for i := 1; i < n; i++ {
		if weights[i] < min {
			min = weights[i]
		}
		if weights[i] > max {
			max = weights[i]
		}
	}

	// Output berat terkecil dan terbesar
	fmt.Printf("\nBerat anak kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat anak kelinci terbesar: %.2f kg\n", max)
}
