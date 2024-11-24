package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel untuk jumlah anak kelinci
	var jumlah int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlah)

	// Validasi jumlah anak kelinci
	if jumlah <= 0 || jumlah > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	// Deklarasi array untuk berat anak kelinci
	beratKelinci := make([]float64, jumlah)
	fmt.Println("Masukkan berat masing-masing anak kelinci:")

	// Input berat anak kelinci
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	// Inisialisasi berat terkecil dan terbesar
	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	// Cari berat terkecil dan terbesar
	for _, berat := range beratKelinci {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
	}

	// Tampilkan hasil
	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}
