package main

import (
	"fmt"
	"strings"
)

// Definisi tipe data untuk array berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rataRata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	// Tampilan header
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("  Program Pengolahan Data Berat Balita")
	fmt.Println(strings.Repeat("=", 40))

	// Meminta pengguna untuk memasukkan jumlah data balita
	fmt.Print("\nMasukan banyak data berat balita: ")
	fmt.Scanln(&n)

	// Memasukkan data berat balita satu per satu
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}

	// Memanggil fungsi untuk menghitung minimum dan maksimum
	hitungMinMax(berat, n, &bMin, &bMax)

	// Menghitung rata-rata berat balita
	rata := rataRata(berat, n)

	// Menampilkan hasil
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("   Hasil Pengolahan Berat Balita")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Berat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita   : %.2f kg\n", rata)
	fmt.Println(strings.Repeat("-", 40))

	// Menutup program dengan ucapan terima kasih
	fmt.Println("\nTerima kasih telah menggunakan program ini!")
	fmt.Println(strings.Repeat("=", 40))
}
