package main

import (
	"fmt"
)

func main() {
	var jumlahWadah, ikanPerWadah int

	// Input jumlah wadah dan jumlah ikan per wadah
	fmt.Print("Masukkan jumlah wadah: ")
	fmt.Scan(&jumlahWadah)

	fmt.Print("Masukkan jumlah ikan per wadah: ")
	fmt.Scan(&ikanPerWadah)

	// Validasi input
	if jumlahWadah <= 0 || jumlahWadah > 1000 || ikanPerWadah <= 0 {
		fmt.Println("Jumlah wadah harus antara 1 hingga 1000, dan jumlah ikan per wadah harus lebih dari 0.")
		return
	}

	// Input berat ikan
	beratIkan := make([]float64, jumlahWadah*ikanPerWadah) // Slice untuk menyimpan berat ikan
	fmt.Println("Masukkan berat setiap ikan (jumlah total: jumlah wadah x jumlah ikan per wadah):")
	for indeks := 0; indeks < jumlahWadah*ikanPerWadah; indeks++ {
		fmt.Printf("Berat ikan ke-%d: ", indeks+1)
		fmt.Scan(&beratIkan[indeks])
	}

	// Menghitung total berat per wadah
	totalBerat := make([]float64, jumlahWadah)
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		for ikan := 0; ikan < ikanPerWadah; ikan++ {
			totalBerat[wadah] += beratIkan[wadah*ikanPerWadah+ikan]
		}
	}

	// Menghitung rata-rata berat per wadah
	var totalKeseluruhanBerat float64
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		totalKeseluruhanBerat += totalBerat[wadah]
	}
	rataRata := totalKeseluruhanBerat / float64(jumlahWadah)

	// Menampilkan hasil
	fmt.Println("\nHasil:")
	fmt.Println("Total berat ikan di setiap wadah:")
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		fmt.Printf("Wadah %d: %.2f\n", wadah+1, totalBerat[wadah])
	}
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRata)
}
