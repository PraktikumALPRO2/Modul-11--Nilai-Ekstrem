package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func hitungRataRata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	fmt.Print("Input banyak data berat balita: ")
	fmt.Scan(&n)

	if n > 100 || n <= 0 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	var dataBerat_145 arrBalita
	for i := 0; i < n; i++ {
		fmt.Printf("input berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat_145[i])
	}

	var bMin, bMax float64
	hitungMinMax(dataBerat_145, n, &bMin, &bMax)
	rataRata := hitungRataRata(dataBerat_145, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
