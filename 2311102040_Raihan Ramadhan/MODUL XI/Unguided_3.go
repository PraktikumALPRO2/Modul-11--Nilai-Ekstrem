package main

import (
	"fmt"
)

// Fungsi untuk menghitung nilai minimum dan maksimum dari array
func hitungMinMax(arrBerat []float64, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for _, berat := range arrBerat {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRatarata(arrBerat []float64) float64 {
	total_2311102040 := 0.0
	for _, berat := range arrBerat {
		total_2311102040 += berat
	}
	return total_2311102040 / float64(len(arrBerat))
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	// Inisialisasi array untuk menyimpan berat balita
	arrBerat := make([]float64, n)

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Deklarasi variabel untuk menyimpan nilai minimum dan maksimum
	var bMin, bMax float64

	// Hitung nilai minimum dan maksimum
	hitungMinMax(arrBerat, &bMin, &bMax)

	// Hitung rata-rata berat balita
	rerata := hitungRatarata(arrBerat)

	// Tampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
