package main

import (
	"fmt"
)

// Mendeklarasikan tipe data struct 'wadah' untuk menampung total berat ikan per wadah
type wadah struct {
	total_beratIkan float64
}

func main() {
	// Mendeklarasikan dua variabel x dan y, untuk jumlah ikan dan kapasitas wadah
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")  // Menampilkan pesan untuk input
	fmt.Scan(&x, &y)  // Membaca input jumlah ikan dan kapasitas wadah

	// Memastikan jumlah ikan yang diinputkan harus antara 1 sampai 1000 dan kapasitas wadah lebih besar dari 0
	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Input tidak valid. Jumlah ikan harus antara 1 dan 1000. Kapasitas wadah harus lebih besar dari 0")  // Menampilkan pesan error
		return  // Menghentikan program jika input tidak valid
	}

	// Array untuk menampung berat ikan, kapasitas maksimal 1000 ikan
	var beratIkan [1000]float64

	// Meminta pengguna untuk memasukkan berat ikan satu per satu
	fmt.Println("Masukkan berat ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])  // Membaca berat ikan dari input pengguna
	}

	// Menghitung jumlah wadah yang diperlukan
	jumlahWadah := (x + y - 1) / y

	// Array untuk menampung wadah, kapasitas maksimal 1000 wadah
	var wadah [1000]wadah

	// Mendistrubusikan berat ikan ke wadah sesuai dengan kapasitas wadah
	for i := 0; i < x; i++ {
		wadahIndex := i / y  // Menentukan indeks wadah berdasarkan indeks ikan
		wadah[wadahIndex].total_beratIkan += beratIkan[i]  // Menambahkan berat ikan ke total berat wadah yang sesuai
	}

	// Menampilkan total berat per wadah
	fmt.Println("Total berat per wadah: ")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, wadah[i].total_beratIkan)  // Menampilkan total berat ikan dalam wadah ke-i
	}

	// Menghitung total berat ikan untuk menghitung rata-rata
	total_beratIkan := 0.0
	for i := 0; i < jumlahWadah; i++ {
		total_beratIkan += wadah[i].total_beratIkan  // Menambahkan total berat wadah ke total keseluruhan
	}

	// Menghitung rata-rata berat per wadah
	rRata := total_beratIkan / float64(jumlahWadah)
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rRata)  // Menampilkan berat rata-rata per wadah
}
