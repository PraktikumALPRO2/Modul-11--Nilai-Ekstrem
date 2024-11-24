// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai minimum dan maksimum
func hitungMinMax(arrBalita []float64) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, berat := range arrBalita {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}
	return min, max
}

// Fungsi untuk menghitung rata-rata
func hitungRataRata(arrBalita []float64) float64 {
	total := 0.0
	for _, berat := range arrBalita {
		total += berat
	}
	return total / float64(len(arrBalita))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	arrBalita := make([]float64, n)

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBalita[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	min, max := hitungMinMax(arrBalita)
	rataRata := hitungRataRata(arrBalita)

	// Tampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
