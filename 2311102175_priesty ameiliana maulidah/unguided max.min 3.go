package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	beratBalita := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	min := beratBalita[0]
	max := beratBalita[0]
	total := 0.0

	for _, berat := range beratBalita {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
		total += berat
	}

	average := total / float64(n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", average)
}
