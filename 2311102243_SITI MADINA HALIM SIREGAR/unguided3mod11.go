package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBalita arrBalita, bMin, bMax *float64) {
	*bMin = arrBalita[0]
	*bMax = arrBalita[0]

	for _, berat := range arrBalita {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

func rerata(arrBalita arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBalita[i]
	}
	return total / float64(n)
}

func main() {
	var x int
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&x)

	var beratBalita arrBalita

	for i := 0; i < x; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	var min, max float64
	hitungMinMax(beratBalita, &min, &max)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	avg := rerata(beratBalita, x)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}
