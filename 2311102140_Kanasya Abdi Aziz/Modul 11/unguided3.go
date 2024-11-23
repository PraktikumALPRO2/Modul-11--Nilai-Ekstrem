package main

import (
	"fmt"
	"math"
)

// Mendefinisikan tipe array untuk berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	*bMin = math.MaxFloat64
	*bMax = -math.MaxFloat64

	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rerata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var beratBalita arrBalita
	var jumlahData int
	var beratMin, beratMax float64

	// Input jumlah data
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&jumlahData)

	// Input data berat balita
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Hitung berat minimum dan maksimum
	hitungMinMax(beratBalita, &beratMin, &beratMax, jumlahData)

	// Hitung rerata
	rerataBalita := rerata(beratBalita, jumlahData)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBalita)
}
