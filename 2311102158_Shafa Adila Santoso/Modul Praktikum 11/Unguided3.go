package main

import "fmt"

func hitungMinMax_158(arrBalita []float64) (float64, float64, float64) {
	bMin := arrBalita[0]
	bMax := arrBalita[0]
	total := 0.0

	for _, berat := range arrBalita {
		if berat < bMin {
			bMin = berat
		}
		if berat > bMax {
			bMax = berat
		}
		total += berat
	}

	rataRata := total / float64(len(arrBalita))
	return bMin, bMax, rataRata
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scanln(&n)

	if n > 0 {
		arrBalita := make([]float64, n)

		fmt.Println("Masukkan data berat balita:")
		for i := 0; i < n; i++ {
			fmt.Printf("Berat balita ke-%d: ", i+1)
			fmt.Scanln(&arrBalita[i])
		}

		bMin, bMax, rataRata := hitungMinMax_158(arrBalita)

		fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
		fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
		fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
	} else {
		fmt.Println("Jumlah balita harus lebih dari 0.")
	}
}
