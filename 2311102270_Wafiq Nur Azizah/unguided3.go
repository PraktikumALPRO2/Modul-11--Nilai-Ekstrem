package main

import "fmt"

type arrBerat [100]float64

func hitungMinMax(arr arrBerat, n int, Min, Max *float64) {
	*Min = arr[0]
	*Max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *Min {
			*Min = arr[i]
		}
		if arr[i] > *Max {
			*Max = arr[i]
		}
	}
}

func rerata(arr arrBerat, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var berat arrBerat
	var n int
	var min, max float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data tidak valid! Harus antara 1 hingga 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)
	rata := rerata(berat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
