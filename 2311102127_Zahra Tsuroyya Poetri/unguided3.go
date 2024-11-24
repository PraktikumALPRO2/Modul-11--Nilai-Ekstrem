package main

import (
	"fmt"
)

// Mendeklarasikan tipe data array 'arrBalita' untuk menampung berat balita, kapasitas maksimal 100
type arrBalita [100]float64

// Fungsi untuk menghitung nilai minimum dan maksimum dari array berat balita
func hitungMinMax(arrBerat arrBalita, n int) (float64, float64) {
	// Inisialisasi nilai minimum dan maksimum 
	bMin := arrBerat[0]
	bMax := arrBerat[0]

	// Melakukan iterasi untuk mencari nilai minimum dan maksimum
	for i := 1; i < n; i++ {
		// Menentukan jika nilai array lebih kecil dari nilai minimum
		if arrBerat[i] < bMin {
			bMin = arrBerat[i]
		}
		// Menentukan jika nilai array lebih besar dari nilai maksimum
		if arrBerat[i] > bMax {
			bMax = arrBerat[i]
		}
	}
	// Mengembalikan nilai minimum dan maksimum
	return bMin, bMax
}

// Fungsi untuk menghitung rata-rata dari array berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	// Menjumlahkan semua nilai dalam array
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	// Mengembalikan nilai rata-rata
	return total / float64(n)
}

func main() {
	var n int               // Deklarasi variabel untuk jumlah data balita
	var berat arrBalita      // Deklarasi array untuk menyimpan berat balita

	// Meminta input jumlah data berat balita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)  

	// Memastikan jumlah data yang diinputkan berada antara 1 dan 100
	if n <= 0 || n > 100 {
		fmt.Println("JInput tidak valid. Jumlah harus antara 1 dan 100")  // Menampilkan pesan jika input tidak valid
		return  // Menghentikan program
	}

	// Meminta input berat balita satu per satu
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)  // Menampilkan pesan untuk input berat balita ke-i
		fmt.Scan(&berat[i])  // Membaca berat balita ke-i
	}

	// Menghitung nilai minimum dan maksimum berat balita
	bMin, bMax := hitungMinMax(berat, n)

	// Menghitung rata-rata berat balita
	rataRata := rerata(berat, n)

	// Menampilkan hasil perhitungan
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)  // Menampilkan berat balita minimum
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)  // Menampilkan berat balita maksimum
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)  // Menampilkan rata-rata berat balita
}
