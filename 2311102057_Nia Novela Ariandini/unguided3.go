package main

import (
	"fmt"
)

//fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat []float64, bMin *float64, bMax *float64) {
	//inisialisasi bMin dan bMax dengan nilai elemen pertama
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for _, berat := range arrBerat {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

//fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arrBerat []float64) float64 {
	var total float64
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	arrBerat := make([]float64, n) //membuat array dengan panjang n

	//input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	//hitung berat minimum, maksimum, dan rata-rata
	var bMin, bMax float64
	hitungMinMax(arrBerat, &bMin, &bMax)
	rerata := hitungRerata(arrBerat)

	//tampilkan hasil
	fmt.Printf("Berat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita : %.2f kg\n", rerata)
}
