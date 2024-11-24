package main

import (
	"fmt"
)

func main() {
	// Input jumlah kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&N)

	// Validasi input
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1-1000")
		return
	}

	// Deklarasi slice untuk menyimpan berat kelinci
	beratKelinci := make([]float64, N)

	// Input berat kelinci dan inisialisasi nilai min dan max
	var beratMin, beratMax float64
	fmt.Println("Masukkan berat masing-masing kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])

		if i == 0 {
			beratMin, beratMax = beratKelinci[i], beratKelinci[i]
		} else {
			if beratKelinci[i] < beratMin {
				beratMin = beratKelinci[i]
			}
			if beratKelinci[i] > beratMax {
				beratMax = beratKelinci[i]
			}
		}
	}

	// Output hasil
	fmt.Printf("\nHasil analisis berat kelinci:\n")
	fmt.Printf("Berat terkecil: %.2f kg\n", beratMin)
	fmt.Printf("Berat terbesar: %.2f kg\n", beratMax)
}
