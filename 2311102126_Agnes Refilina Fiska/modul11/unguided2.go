package main

import (
	"fmt"
)

func main() {
	var x, y int
	var weights = make([]float64, 1000)      // Array untuk berat ikan
	var containers = make([][]float64, 1000) // Array 2D untuk wadah

	// Input jumlah ikan (x)
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scan(&x)

	// Input jumlah wadah (y)
	fmt.Print("Masukkan jumlah wadah yang tersedia (y): ")
	fmt.Scan(&y)

	// Input berat untuk setiap ikan
	fmt.Println("\nMasukkan berat untuk setiap ikan (kg):")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Inisialisasi array wadah
	for i := 0; i < y; i++ {
		containers[i] = make([]float64, 0)
	}

	// Distribusi ikan ke wadah
	currentContainer := 0
	for i := 0; i < x; i++ {
		containers[currentContainer] = append(containers[currentContainer], weights[i])
		currentContainer = (currentContainer + 1) % y
	}

	// Output total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i := 0; i < y; i++ {
		var totalWeight float64
		for _, weight := range containers[i] {
			totalWeight += weight
		}
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, totalWeight)
	}

	// Hitung dan output rata-rata berat ikan di setiap wadah
	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i := 0; i < y; i++ {
		var totalWeight float64
		numFish := len(containers[i])
		for _, weight := range containers[i] {
			totalWeight += weight
		}
		if numFish > 0 {
			average := totalWeight / float64(numFish)
			fmt.Printf("Wadah %d: %.2f kg\n", i+1, average)
		} else {
			fmt.Printf("Wadah %d: 0.00 kg (kosong)\n", i+1)
		}
	}
}
