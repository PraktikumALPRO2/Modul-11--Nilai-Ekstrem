package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var n int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	// Deklarasi array untuk menampung berat kelinci
	berat := make([]float64, n)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Inisialisasi min dan max
	min := berat[0]
	max := berat[0]

	// Mencari min dan max
	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}