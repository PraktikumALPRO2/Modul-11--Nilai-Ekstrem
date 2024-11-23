package main

import (
	"fmt"
)

func main() {
	// Deklarasi array dengan kapasitas 1000
	var beratIkan [1000]float64
	var beratWadah [1000]float64

	// Input jumlah ikan (x)
	var x int
	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&x)

	// Input kapasitas wadah (y)
	var y int
	fmt.Print("Masukkan kapasitas wadah (y): ")
	fmt.Scan(&y)

	// Input berat ikan
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Hitung jumlah wadah yang diperlukan
	jumWadah := (x + y - 1) / y

	// Hitung total berat per wadah
	for i := 0; i < x; i++ {
		wadahKe := i / y
		beratWadah[wadahKe] += beratIkan[i]
	}

	// Output berat total di setiap wadah
	fmt.Println("Berat total di setiap wadah:")
	for i := 0; i < jumWadah; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, beratWadah[i])
	}

	// Hitung dan output rata-rata
	var totalBerat float64
	for i := 0; i < jumWadah; i++ {
		totalBerat += beratWadah[i]
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", totalBerat/float64(jumWadah))
}
