package main

import (
	"fmt"
)

// Definisi tipe data untuk array berat balita
type arrayBerat [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func cariMinMax(berat arrayBerat, jumlah int, min *float64, max *float64) {
	*min = berat[0]
	*max = berat[0]

	for i := 1; i < jumlah; i++ {
		if berat[i] < *min {
			*min = berat[i]
		}
		if berat[i] > *max {
			*max = berat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRataRata(berat arrayBerat, jumlah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlah; i++ {
		total += berat[i]
	}
	return total / float64(jumlah)
}

// Fungsi utama
func main() {
	var jumlahBalita int
	var beratBalita arrayBerat

	// Input jumlah balita
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahBalita)

	// Validasi jumlah balita
	if jumlahBalita <= 0 || jumlahBalita > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	fmt.Println("Masukkan berat balita:")
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Deklarasi variabel untuk hasil
	var beratMin, beratMax float64

	// Panggil fungsi cariMinMax
	cariMinMax(beratBalita, jumlahBalita, &beratMin, &beratMax)

	// Panggil fungsi hitungRataRata
	rataRata := hitungRataRata(beratBalita, jumlahBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
