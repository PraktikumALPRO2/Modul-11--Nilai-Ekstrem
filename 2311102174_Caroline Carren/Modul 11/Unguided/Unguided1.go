// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("    Program Pengolahan Berat Kelinci")
		fmt.Println(strings.Repeat("=", 40))

		// Input: Jumlah kelinci
		var n int
		fmt.Print("\nMasukkan jumlah kelinci: ")
		fmt.Scanln(&n)

		// Memastikan jumlah kelinci lebih dari 0
		if n <= 0 {
			fmt.Println("\n[!] Jumlah kelinci harus lebih dari 0. Silakan coba lagi.")
			continue
		}

		// Membuat slice untuk menyimpan berat kelinci
		weights := make([]float64, n)
		fmt.Println("\nMasukkan berat masing-masing kelinci:")
		for i := 0; i < n; i++ {
			fmt.Printf("  - Berat kelinci %d: ", i+1)
			fmt.Scanln(&weights[i])
		}

		// Mengurutkan berat kelinci dalam urutan menaik
		sort.Float64s(weights)

		// Output: Berat kelinci terkecil dan terbesar
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("        Hasil Pengolahan")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("🐇 Berat terkecil : %.2f kg\n", weights[0])
		fmt.Printf("🐇 Berat terbesar : %.2f kg\n", weights[n-1])
		fmt.Println("\n📋 Daftar berat kelinci (terurut):")
		for i, weight := range weights {
			fmt.Printf("  %2d. %.2f kg\n", i+1, weight)
		}
		fmt.Println(strings.Repeat("-", 40))

		// Menanyakan apakah pengguna ingin mengulangi
		var again string
		fmt.Print("\n🔄 Apakah Anda ingin memasukkan data lagi? (ya/tidak): ")
		fmt.Scanln(&again)

		// Jika pengguna tidak ingin mengulangi, keluar dari loop
		if strings.ToLower(again) != "ya" {
			fmt.Println("\n✨ Terima kasih telah menggunakan program ini. ✨")
			break
		}
	}
}
