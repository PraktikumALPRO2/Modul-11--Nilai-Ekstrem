package main

import (
	"fmt"
)

func main() {
	var beratKelinci []float64

	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1-1000")
		return
	}

	beratKelinci = make([]float64, jumlahKelinci)

	fmt.Println("Masukkan berat masing-masing kelinci:")
	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratMin, beratMax := beratKelinci[0], beratKelinci[0]

	for _, berat := range beratKelinci {
		if berat < beratMin {
			beratMin = berat
		}
		if berat > beratMax {
			beratMax = berat
		}
	}

	fmt.Printf("\nHasil analisis berat kelinci:\n")
	fmt.Printf("Berat terkecil: %.2f kg\n", beratMin)
	fmt.Printf("Berat terbesar: %.2f kg\n", beratMax)
}
