package main

import (
	"fmt"
	"math"
)

// Fungsi untuk meminta input berat kelinci
func inputBeratKelinci(N235 int) []float64 {
	berat235 := make([]float64, N235)
	fmt.Println("Masukkan berat anak kelinci satu per satu:")
	for i := 0; i < N235; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat235[i])
	}
	return berat235
}

// Fungsi untuk mencari berat terkecil dan terbesar
func cariBeratTerbesarTerkecil(berat235 []float64) (float64, float64) {
	minBerat235 := math.MaxFloat64
	maxBerat235 := -math.MaxFloat64

	for _, b235 := range berat235 {
		if b235 < minBerat235 {
			minBerat235 = b235
		}
		if b235 > maxBerat235 {
			maxBerat235 = b235
		}
	}
	return minBerat235, maxBerat235
}

// Fungsi utama
func main() {
	// Input jumlah anak kelinci
	var N235 int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N235)

	// Validasi jumlah kelinci
	if N235 <= 0 || N235 > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	// Input berat kelinci
	berat235 := inputBeratKelinci(N235)

	// Cari berat terkecil dan terbesar
	minBerat235, maxBerat235 := cariBeratTerbesarTerkecil(berat235)

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", minBerat235)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat235)
}
