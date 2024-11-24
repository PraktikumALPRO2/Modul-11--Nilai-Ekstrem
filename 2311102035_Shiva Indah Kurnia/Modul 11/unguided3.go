package main

import (
	"fmt"
	"math"
)

func main() {
	// Input jumlah data
	var jumlahData int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahData)

	// Validasi input
	if jumlahData <= 0 || jumlahData > 100 {
		fmt.Println("Jumlah data harus antara 1-100")
		return
	}

	// Deklarasi slice untuk menyimpan berat balita
	beratBalita := make([]float64, jumlahData)

	// Input data berat balita
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Inisialisasi nilai min, max, dan total
	beratMin := math.MaxFloat64
	beratMax := -math.MaxFloat64
	var total float64

	// Hitung berat minimum, maksimum, dan total
	for _, berat := range beratBalita {
		if berat < beratMin {
			beratMin = berat
		}
		if berat > beratMax {
			beratMax = berat
		}
		total += berat
	}

	// Hitung rerata
	rerataBalita := total / float64(jumlahData)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBalita)
}
