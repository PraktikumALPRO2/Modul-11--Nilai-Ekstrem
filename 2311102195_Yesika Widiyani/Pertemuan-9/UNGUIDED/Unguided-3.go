package main

import (
	"fmt"
)

type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rata-rata berat balita
func hitungRataRata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	// Input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data tidak valid. Harus antara 1 dan 100.")
		return
	}

	// Input data berat balita
	fmt.Println("Masukkan berat balita (dalam kg):")
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung minimum dan maksimum
	hitungMinMax(berat, n, &bMin, &bMax)

	// Hitung rata-rata
	rataRata := hitungRataRata(berat, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
