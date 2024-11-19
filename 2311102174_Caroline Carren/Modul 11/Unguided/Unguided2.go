// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input: Masukan
	var x, y int
	var ikan []float64

	// Meminta pengguna untuk memasukkan jumlah ikan
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("    Program Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("\nMasukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scanln(&x)
	if x <= 0 {
		fmt.Println("\n[!] Jumlah ikan harus positif.")
		return // Keluar dari program jika jumlah ikan tidak valid
	}

	// Meminta pengguna untuk memasukkan jumlah wadah
	fmt.Print("\nMasukkan jumlah wadah (y): ")
	fmt.Scanln(&y)
	if y <= 0 {
		fmt.Println("\n[!] Jumlah wadah harus positif.")
		return // Keluar dari program jika jumlah wadah tidak valid
	}

	// Meminta pengguna untuk memasukkan berat setiap ikan satu per satu
	fmt.Println("\nMasukkan berat setiap ikan (dalam kg):")
	for i := 0; i < x; i++ {
		var berat float64
		fmt.Printf("  - Berat ikan %d: ", i+1)
		_, err := fmt.Scanln(&berat)
		if err != nil {
			fmt.Println("\n[!] Input berat ikan tidak valid.")
			return // Keluar dari program jika ada kesalahan input
		}
		ikan = append(ikan, berat)
	}

	// Hitung total berat ikan di setiap wadah
	wadah := make([]float64, y) // Membuat slice untuk menyimpan total berat di setiap wadah
	for i := 0; i < x; i++ {
		wadah[i%y] += ikan[i] // Menambahkan berat ikan ke wadah yang sesuai
	}

	// Hitung rata-rata berat ikan di setiap wadah
	rataRata := 0.0
	for _, v := range wadah {
		rataRata += v // Menjumlahkan total berat di semua wadah
	}
	rataRata /= float64(x) // Rata-rata berdasarkan jumlah ikan

	// Output: Keluaran
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     Hasil Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("-", 40))

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, v := range wadah {
		fmt.Printf("  Wadah %d: %.2f kg\n", i+1, v) // Menampilkan total berat ikan di setiap wadah
	}

	// Menampilkan berat rata-rata ikan di setiap wadah
	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	fmt.Printf("  %.2f kg\n", rataRata) // Menampilkan rata-rata berat ikan

	// Akhir dari program
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Terima kasih telah menggunakan program ini!")
	fmt.Println(strings.Repeat("=", 40))
}
