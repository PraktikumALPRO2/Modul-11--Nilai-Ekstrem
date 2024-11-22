package main

import (
	"fmt"
	"math"
)

// Fungsi untuk membaca berat ikan dari input pengguna
func bacaBeratIkan235(jumlahIkan235 int) []float64 {
	beratIkan235 := make([]float64, jumlahIkan235)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < jumlahIkan235; i++ {
		fmt.Scan(&beratIkan235[i])
	}
	return beratIkan235
}

// Fungsi untuk mendistribusikan ikan ke dalam wadah
func distribusikanIkan235(beratIkan235 []float64, kapasitasWadah235 int) [][]float64 {
	jumlahWadah235 := int(math.Ceil(float64(len(beratIkan235)) / float64(kapasitasWadah235)))
	wadah235 := make([][]float64, jumlahWadah235)

	for i, berat235 := range beratIkan235 {
		indeksWadah235 := i / kapasitasWadah235
		wadah235[indeksWadah235] = append(wadah235[indeksWadah235], berat235)
	}
	return wadah235
}

// Fungsi untuk menghitung total berat ikan di setiap wadah
func hitungTotalBerat235(wadah235 [][]float64) []float64 {
	totalBerat235 := make([]float64, len(wadah235))
	for i, isiWadah235 := range wadah235 {
		var jumlah235 float64
		for _, berat235 := range isiWadah235 {
			jumlah235 += berat235
		}
		totalBerat235[i] = jumlah235
	}
	return totalBerat235
}

// Fungsi untuk menghitung rata-rata berat ikan di setiap wadah
func hitungRataRataBerat235(wadah235 [][]float64) []float64 {
	rataRataBerat235 := make([]float64, len(wadah235))
	for i, isiWadah235 := range wadah235 {
		var jumlah235 float64
		for _, berat235 := range isiWadah235 {
			jumlah235 += berat235
		}
		rataRataBerat235[i] = jumlah235 / float64(len(isiWadah235))
	}
	return rataRataBerat235
}

// Fungsi untuk menampilkan hasil total dan rata-rata berat
func tampilkanHasil235(totalBerat235, rataRataBerat235 []float64) {
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total235 := range totalBerat235 {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total235)
	}

	fmt.Println("\nRata-rata berat ikan di setiap wadah:")
	for i, rataRata235 := range rataRataBerat235 {
		fmt.Printf("Wadah %d: %.2f\n", i+1, rataRata235)
	}
}

func main() {
	var jumlahIkan235, kapasitasWadah235 int
	fmt.Println("Masukkan jumlah ikan yang akan dijual (x) dan kapasitas wadah (y):")
	fmt.Scan(&jumlahIkan235, &kapasitasWadah235)

	// Memanggil fungsi untuk membaca input dan memproses data
	beratIkan235 := bacaBeratIkan235(jumlahIkan235)
	wadah235 := distribusikanIkan235(beratIkan235, kapasitasWadah235)
	totalBerat235 := hitungTotalBerat235(wadah235)
	rataRataBerat235 := hitungRataRataBerat235(wadah235)

	// Menampilkan hasil
	tampilkanHasil235(totalBerat235, rataRataBerat235)
}
