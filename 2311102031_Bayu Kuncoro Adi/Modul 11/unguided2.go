package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var x_2311102031, y_2311102031 int
	var weights_2311102031 []float64

	// Input jumlah ikan dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x_2311102031, &y_2311102031)

	// Validasi jumlah ikan dan kapasitas wadah
	if x_2311102031 <= 0 || y_2311102031 <= 0 || x_2311102031 > 1000 {
		fmt.Println("Jumlah ikan (x) harus antara 1 dan 1000, dan kapasitas wadah (y) harus lebih dari 0.")
		return
	}

	// Input berat ikan
	weights_2311102031 = make([]float64, x_2311102031)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x_2311102031; i++ {
		fmt.Scan(&weights_2311102031[i])
	}

	// Hitung total berat per wadah
	var totalWeights_2311102031 []float64
	currentWeight_2311102031 := 0.0

	for i, weight := range weights_2311102031 {
		currentWeight_2311102031 += weight

		// Jika wadah penuh atau ikan terakhir
		if (i+1)%y_2311102031 == 0 || i == x_2311102031-1 {
			totalWeights_2311102031 = append(totalWeights_2311102031, currentWeight_2311102031)
			currentWeight_2311102031 = 0.0
		}
	}

	// Hitung rata-rata berat per wadah
	totalSum_2311102031 := 0.0
	for _, total := range totalWeights_2311102031 {
		totalSum_2311102031 += total
	}
	averageWeight_2311102031 := totalSum_2311102031 / float64(len(totalWeights_2311102031))

	// Output hasil
	fmt.Println("Total berat ikan di setiap wadah:")
	for _, total := range totalWeights_2311102031 {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", averageWeight_2311102031)
}
