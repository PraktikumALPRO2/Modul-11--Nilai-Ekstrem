package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Masukan tidak valid. Pastikan x > 0, y > 0, dan x <= 1000.")
		return
	}

	var weights [1000]float64

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	var containerWeights []float64
	var currentWeight float64

	for i := 0; i < x; i++ {
		currentWeight += weights[i]
		if (i+1)%y == 0 || i == x-1 {
			containerWeights = append(containerWeights, currentWeight)
			currentWeight = 0
		}
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, weight := range containerWeights {
		fmt.Printf("%.2f ", weight)
	}
	fmt.Println()

	totalWeight := 0.0
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(len(containerWeights))

	fmt.Printf("Berat rata-rata ikan per wadah: %.2f\n", averageWeight)
}
