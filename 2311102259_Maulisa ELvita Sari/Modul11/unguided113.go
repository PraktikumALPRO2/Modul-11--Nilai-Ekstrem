package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, N int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < N; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func hitungRata(arrBerat arrBalita, N int) float64 {
	var total float64 = 0
	for i := 0; i < N; i++ {
		total += arrBerat[i]
	}
	return total / float64(N)
}

func main() {
	var arrBerat arrBalita
	var N int

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&N)

	if N <= 0 || N > 100 {
		fmt.Println("Jumlah data balita tidak valid. Harus antara 1 hingga 100.")
		return
	}

	fmt.Println("Masukkan berat balita:")
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	var bMin, bMax float64

	hitungMinMax(arrBerat, N, &bMin, &bMax)

	rataRata := hitungRata(arrBerat, N)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
