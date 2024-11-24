package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Println("Masukkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y):")
	fmt.Scan(&x, &y)

	// Membuat array untuk menyimpan berat ikan
	beratIkan_2311102040 := make([]float64, x)
	fmt.Println("Masukkan berat ikan satu per satu:")

	// Memasukkan berat ikan ke dalam array
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan_2311102040[i])
	}

	// Menghitung jumlah wadah
	jumlahWadah := (x + y - 1) / y // Menggunakan pembulatan ke atas

	// Menyimpan total berat di setiap wadah
	totalBeratPerWadah := make([]float64, jumlahWadah)
	for i := 0; i < x; i++ {
		wadahIndex := i / y
		totalBeratPerWadah[wadahIndex] += beratIkan_2311102040[i]
	}

	// Menampilkan total berat per wadah
	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, totalBeratPerWadah[i])
	}

	// Menghitung berat rata-rata
	totalBeratSemuaWadah := 0.0
	for _, berat := range totalBeratPerWadah {
		totalBeratSemuaWadah += berat
	}
	beratRataRata := totalBeratSemuaWadah / float64(jumlahWadah)

	// Menampilkan berat rata-rata
	fmt.Printf("Berat rata-rata di setiap wadah: %.2f\n", beratRataRata)
}
