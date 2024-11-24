package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 || y > x {
		fmt.Println("Input tidak valid. Pastikan x dan y berada dalam batas yang benar.")
		return
	}
	weights := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	totalWeights := make([]float64, (x+y-1)/y) 

	for i := 0; i < x; i++ {
		wadahIndex := i / y
		totalWeights[wadahIndex] += weights[i]
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, totalWeight := range totalWeights {
		fmt.Printf("%.2f ", totalWeight)
	}
	fmt.Println()

	totalWeightAllWadah := 0.0
	for _, totalWeight := range totalWeights {
		totalWeightAllWadah += totalWeight
	}

	averageWeight := totalWeightAllWadah / float64(len(totalWeights))
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", averageWeight)
}
