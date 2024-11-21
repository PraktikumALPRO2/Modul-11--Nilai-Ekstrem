package main

import (
	"fmt"
	"math"
)

func main() {
	// Input jumlah ikan dan kapasitas wadah
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas wadah (y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || y <= 0 || y > x {
		fmt.Println("Input tidak valid. Pastikan x > 0, y > 0, dan y <= x.")
		return
	}

	// Input berat ikan
	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Hitung jumlah wadah
	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))
	totalWadah := make([]float64, jumlahWadah)

	// Distribusi berat ikan ke wadah
	for i := 0; i < x; i++ {
		wadahKe := i / y
		totalWadah[wadahKe] += beratIkan[i]
	}

	// Hitung rata-rata berat per wadah
	totalBerat := 0.0
	for _, berat := range totalWadah {
		totalBerat += berat
	}
	rataRata := totalBerat / float64(jumlahWadah)

	// Output hasil
	fmt.Println("Berat total di setiap wadah:")
	for i, berat := range totalWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", rataRata)
}