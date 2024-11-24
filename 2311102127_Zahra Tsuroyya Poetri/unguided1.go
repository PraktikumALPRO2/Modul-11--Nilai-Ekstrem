package main

import (
	"fmt"
)

// Mendeklrasikan tipe data beratKelinci dengan array berkapasitas 1000
type beratKelinci [1000]float64

func main() {
	// Mendeklarasikan variabel n untuk menyimpan jumlah kelinci
	var n int
	// Mendeklarasikan variabel berat bertipe array beratKelinci untuk menyimpan berat kelinci
	var berat beratKelinci

	// Menampilkan pesan untuk meminta input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	// Memastikan jumlah anak kelinci yang diinputkan harus di antara 1 dan 1000
	if n <= 0 || n > 1000 {
		// Jika jumlah kelinci tidak valid, tampilkan pesan error dan hentikan eksekusi program
		fmt.Println("Input tidak valid. Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	// Menampilkan pesan untuk meminta input berat anak kelinci
	fmt.Println("Masukkan berat anak kelinci:")
	// Loop untuk menerima input berat kelinci sesuai dengan jumlah n
	for i := 0; i < n; i++ {
		// Menampilkan pesan untuk meminta berat anak kelinci ke-i
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		// Menerima input berat anak kelinci ke-i
		fmt.Scan(&berat[i])
	}

	// Inisialisasi nilai awal untuk pencarian berat terkecil (min) dan terbesar (max)
	minBerat := berat[0]
	maxBerat := berat[0]

	// Melakukan iterasi untuk mencari berat terkecil dan terbesar
	for i := 1; i < n; i++ {
		// Jika berat kelinci ke-i lebih kecil dari minBerat, perbarui minBerat
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		// Jika berat kelinci ke-i lebih besar dari maxBerat, perbarui maxBerat
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	// Menampilkan hasil berat terkecil dan terbesar 
	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
