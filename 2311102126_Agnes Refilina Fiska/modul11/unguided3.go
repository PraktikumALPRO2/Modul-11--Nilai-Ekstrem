package main

import (
	"fmt"
)

// Mendefinisikan tipe array untuk berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Mencari nilai minimum dan maksimum
	for i := 1; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
}

// Fungsi untuk menghitung rerata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var dataBerat arrBalita
	var jumlahData int
	var min, max float64

	// Input jumlah data
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&jumlahData)

	// Input berat balita
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	// Hitung min dan max
	hitungMinMax(dataBerat, &min, &max)

	// Hitung rerata
	rata := rerata(dataBerat, jumlahData)

	// Output hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
