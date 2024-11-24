// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Input jumlah wadah (x) dan jumlah ikan per wadah (y)
	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Jumlah wadah dan ikan per wadah harus lebih dari 0 dan jumlah wadah maksimal 1000.")
		return
	}

	// Input berat ikan
	berat := make([]float64, x*y) // Array untuk menampung berat ikan
	fmt.Println("Masukkan berat ikan (sejumlah x * y):")
	for i := 0; i < x*y; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung total berat per wadah
	totalBeratPerWadah := make([]float64, x)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			totalBeratPerWadah[i] += berat[i*y+j]
		}
	}

	// Hitung berat rata-rata per wadah
	var totalBeratSemua float64
	for i := 0; i < x; i++ {
		totalBeratSemua += totalBeratPerWadah[i]
	}
	rataRataBerat := totalBeratSemua / float64(x)

	// Output hasil
	fmt.Println("\nHasil:")
	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, totalBeratPerWadah[i])
	}
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRataBerat)
}
