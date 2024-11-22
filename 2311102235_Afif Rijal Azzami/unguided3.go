package main

import "fmt"

// Tipe data array untuk menyimpan berat balita
type arrBalita235 [100]float64

// Fungsi untuk menghitung berat min dan max
func hitungMinMax235(arrBerat235 arrBalita235, n235 int) (minBerat235, maxBerat235 float64) {
	if n235 == 0 {
		return 0, 0
	}

	// untuk nilai minimum dan maksimum
	minBerat235, maxBerat235 = arrBerat235[0], arrBerat235[0]
	for i235 := 1; i235 < n235; i235++ {
		if arrBerat235[i235] < minBerat235 {
			minBerat235 = arrBerat235[i235]
		}
		if arrBerat235[i235] > maxBerat235 {
			maxBerat235 = arrBerat235[i235]
		}
	}
	return
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata235(arrBerat235 arrBalita235, n235 int) float64 {
	if n235 == 0 {
		return 0
	}

	var total235 float64
	for i235 := 0; i235 < n235; i235++ {
		total235 += arrBerat235[i235]
	}
	return total235 / float64(n235)
}

func main() {
	var n235 int
	var dataBerat235 arrBalita235

	for i235 := 0; i235 < len(dataBerat235); i235++ {
		dataBerat235[i235] = 0
	}

	// Input jumlah balita yang ingin di data
	fmt.Print("Banyak data berat balita : ")
	fmt.Scan(&n235)

	// mevalidasi jumlah data
	if n235 <= 0 || n235 > 100 {
		fmt.Println("Jumlah data tidak lebih dari 100.")
		return
	}

	// Input data berat balita
	for i235 := 0; i235 < n235; i235++ {
		fmt.Printf("Berat balita ke - %d : ", i235+1)
		fmt.Scan(&dataBerat235[i235])
	}

	// untuk menghitung minimum, maksimum, dan rata-rata berat
	minBerat235, maxBerat235 := hitungMinMax235(dataBerat235, n235)
	rataRata235 := hitungRerata235(dataBerat235, n235)

	// Output akhir
	fmt.Printf("\nBerat Balita Minimum : %.2f kg\n", minBerat235)
	fmt.Printf("Berat Balita Maksimal : %.2f kg\n", maxBerat235)
	fmt.Printf("Berat Balita Rata-Rata : %.2f kg\n", rataRata235)
}
