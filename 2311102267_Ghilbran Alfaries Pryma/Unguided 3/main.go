package main

import (
	"fmt"
)

// Tipe data array untuk berat balita
type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat []float64) (float64, float64) {
	bMin := arrBerat[0]
	bMax := arrBerat[0]

	for _, berat := range arrBerat {
		if berat < bMin {
			bMin = berat
		}
		if berat > bMax {
			bMax = berat
		}
	}

	return bMin, bMax
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat []float64) float64 {
	total := 0.0
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

// Program utama
func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	var beratBalita arrBalita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Konversi array ke slice
	dataBalita := beratBalita[:n]

	// Hitung berat minimum dan maksimum
	bMin, bMax := hitungMinMax(dataBalita)

	// Hitung rerata
	rataRata := rerata(dataBalita)

	// Tampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
