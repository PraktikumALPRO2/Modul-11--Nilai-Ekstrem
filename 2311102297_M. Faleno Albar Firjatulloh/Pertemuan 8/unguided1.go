package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi array dengan kapasitas 1000
	var beratKelinci [1000]float64

	// Input jumlah kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&N)

	// Validasi input
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1-1000")
		return
	}

	// Input berat kelinci
	fmt.Println("Masukkan berat masing-masing kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	// Inisialisasi nilai minimum dan maksimum
	beratMin := math.MaxFloat64
	beratMax := -math.MaxFloat64

	// Mencari berat minimum dan maksimum
	for i := 0; i < N; i++ {
		// Update berat minimum
		if beratKelinci[i] < beratMin {
			beratMin = beratKelinci[i]
		}

		// Update berat maksimum
		if beratKelinci[i] > beratMax {
			beratMax = beratKelinci[i]
		}
	}

	// Output hasil
	fmt.Printf("\nHasil analisis berat kelinci:\n")
	fmt.Printf("Berat terkecil: %.2f kg\n", beratMin)
	fmt.Printf("Berat terbesar: %.2f kg\n", beratMax)
}
