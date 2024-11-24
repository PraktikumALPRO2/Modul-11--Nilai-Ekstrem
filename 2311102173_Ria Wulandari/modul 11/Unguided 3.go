package main

import (
	"fmt"
	"math"
)

// Fungsi untuk mencari nilai ekstrem (minimum dan maksimum)
func cariEkstrem(beratBalita []float64) (float64, float64) {
	beratMinimal := math.MaxFloat64
	beratMaksimal := -math.MaxFloat64

	for _, berat := range beratBalita {
		if berat < beratMinimal {
			beratMinimal = berat
		}
		if berat > beratMaksimal {
			beratMaksimal = berat
		}
	}
	return beratMinimal, beratMaksimal
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRata(beratBalita []float64) float64 {
	totalBerat := 0.0
	for _, berat := range beratBalita {
		totalBerat += berat
	}
	return totalBerat / float64(len(beratBalita))
}

func main() {
	var jumlahData int
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahData)

	// Validasi jumlah data
	if jumlahData <= 0 || jumlahData > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	beratBalita := make([]float64, jumlahData)

	// Input berat balita
	for indeks := 0; indeks < jumlahData; indeks++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", indeks+1)
		fmt.Scan(&beratBalita[indeks])
	}

	// Menghitung nilai ekstrem dan rata-rata
	beratMin, beratMax := cariEkstrem(beratBalita)
	rataBerat := hitungRata(beratBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita terendah: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita tertinggi: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataBerat)
}
