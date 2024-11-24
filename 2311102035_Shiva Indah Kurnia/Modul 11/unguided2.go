package main

import (
	"fmt"
)

func main() {
	// Input jumlah ikan (x) dan kapasitas wadah (y)
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas wadah (y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0")
		return
	}

	// Deklarasi slice untuk berat ikan dan wadah
	beratIkan := make([]float64, x)
	jumWadah := (x + y - 1) / y // Hitung jumlah wadah
	beratWadah := make([]float64, jumWadah)

	// Input berat ikan
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])

		// Distribusi berat ikan ke wadah
		wadahKe := i / y
		beratWadah[wadahKe] += beratIkan[i]
	}

	// Output berat total di setiap wadah
	fmt.Println("Berat total di setiap wadah:")
	for i, berat := range beratWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}

	// Hitung rata-rata berat ikan per wadah
	var totalBerat float64
	for _, berat := range beratWadah {
		totalBerat += berat
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", totalBerat/float64(jumWadah))
}
