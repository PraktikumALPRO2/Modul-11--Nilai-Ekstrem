// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	// Cek apakah jumlah anak kelinci valid
	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	berat := make([]float64, n) // Array untuk menampung berat anak kelinci
	fmt.Println("Masukkan berat anak kelinci:")

	// Inisialisasi nilai ekstrem
	minBerat := math.MaxFloat64
	maxBerat := -math.MaxFloat64

	// Input berat dan cari nilai minimum dan maksimum
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		// Update nilai minimum dan maksimum
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	// Output hasil
	fmt.Printf("Berat kelinci terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxBerat)
}
