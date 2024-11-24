package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&jumlahKelinci)

	// Validasi jumlah anak kelinci
	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	beratKelinci := make([]float64, jumlahKelinci) // Slice untuk berat kelinci
	fmt.Println("Masukkan berat setiap anak kelinci:")

	// Nilai ekstrem untuk berat
	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	// Loop untuk input berat dan cari nilai ekstrem
	for indeks := 0; indeks < jumlahKelinci; indeks++ {
		fmt.Printf("Berat kelinci ke-%d: ", indeks+1)
		fmt.Scan(&beratKelinci[indeks])

		// Pembaruan nilai ekstrem
		if beratKelinci[indeks] < beratTerkecil {
			beratTerkecil = beratKelinci[indeks]
		}
		if beratKelinci[indeks] > beratTerbesar {
			beratTerbesar = beratKelinci[indeks]
		}
	}

	// Menampilkan hasil
	fmt.Printf("Berat kelinci terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", beratTerbesar)
}
