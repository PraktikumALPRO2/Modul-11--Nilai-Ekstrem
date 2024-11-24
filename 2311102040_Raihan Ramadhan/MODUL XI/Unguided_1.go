package main

import (
	"fmt"
)

func main() {
	const kapasitas_2311102040 = 1000
	var berat [kapasitas_2311102040]float64 // Array dengan kapasitas tetap 1000

	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	// Validasi jumlah anak kelinci
	if n <= 0 || n > kapasitas_2311102040  {
		fmt.Printf("Jumlah anak kelinci harus antara 1 dan %d.\n", kapasitas_2311102040 )
		return
	}

	fmt.Println("Masukkan berat anak kelinci:")

	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Mencari berat terkecil dan terbesar
	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
