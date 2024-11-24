package main

import (
	"fmt"
	"math"
)

func hitungStatistik(arrBerat []float64) (float64, float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64
	total := 0.0

	for _, berat := range arrBerat {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
		total += berat
	}

	rerata := total / float64(len(arrBerat))
	return min, max, rerata
}

func main() {
	var jumlahData int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&jumlahData)

	beratBalita := make([]float64, jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	beratMin, beratMax, rerataBalita := hitungStatistik(beratBalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBalita)
}
