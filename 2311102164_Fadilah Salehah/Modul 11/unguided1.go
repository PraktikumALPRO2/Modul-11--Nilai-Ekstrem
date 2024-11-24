// Fadilah Salehah 
// 2311102164 
// S1 IF 11 05 

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("   Aplikasi Analisis Berat Kelinci")
		fmt.Println(strings.Repeat("=", 40))

		// Input: Jumlah kelinci
		var jumlah int
		fmt.Print("\nMasukkan jumlah kelinci: ")
		fmt.Scanln(&jumlah)

		// Validasi: Jumlah kelinci harus lebih dari 0
		if jumlah <= 0 {
			fmt.Println("\n[!] Jumlah kelinci tidak boleh kurang dari 1. Silakan coba lagi.")
			continue
		}

		// Membuat slice untuk berat kelinci
		beratKelinci := make([]float64, jumlah)
		fmt.Println("\nMasukkan berat kelinci satu per satu:")
		for i := 0; i < jumlah; i++ {
			fmt.Printf("  - Berat kelinci ke-%d: ", i+1)
			fmt.Scanln(&beratKelinci[i])
		}

		// Mengurutkan berat kelinci secara menaik
		sort.Float64s(beratKelinci)

		// Output: Menampilkan berat terkecil dan terbesar
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("         Ringkasan Data Berat")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("🐇 Berat terkecil : %.2f kg\n", beratKelinci[0])
		fmt.Printf("🐇 Berat terbesar : %.2f kg\n", beratKelinci[jumlah-1])
		fmt.Println("\n📋 Daftar berat kelinci (urut):")
		for i, berat := range beratKelinci {
			fmt.Printf("  %2d. %.2f kg\n", i+1, berat)
		}
		fmt.Println(strings.Repeat("-", 40))

		// Menanyakan apakah program akan diulangi
		var ulangi string
		fmt.Print("\n🔄 Ingin mengolah data lagi? (ya/tidak): ")
		fmt.Scanln(&ulangi)

		// Keluar jika jawaban bukan "ya"
		if strings.ToLower(ulangi) != "ya" {
			fmt.Println("\n✨ Terima kasih telah menggunakan aplikasi ini. ✨")
			break
		}
	}
}
