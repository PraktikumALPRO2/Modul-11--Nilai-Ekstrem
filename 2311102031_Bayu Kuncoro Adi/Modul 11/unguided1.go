package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var n_2311102031 int
	var weights_2311102031 []float64

	// Input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n_2311102031)

	// Validasi jumlah anak kelinci
	if n_2311102031 <= 0 || n_2311102031 > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	// Input berat anak kelinci
	weights_2311102031 = make([]float64, n_2311102031)
	fmt.Println("Masukkan berat masing-masing anak kelinci:")
	for i_2311102031 := 0; i_2311102031 < n_2311102031; i_2311102031++ {
		fmt.Scan(&weights_2311102031[i_2311102031])
	}

	// Inisialisasi min dan max
	minWeight_2311102031 := weights_2311102031[0]
	maxWeight_2311102031 := weights_2311102031[0]

	// Cari berat terkecil dan terbesar
	for _, weight_2311102031 := range weights_2311102031 {
		if weight_2311102031 < minWeight_2311102031 {
			minWeight_2311102031 = weight_2311102031
		}
		if weight_2311102031 > maxWeight_2311102031 {
			maxWeight_2311102031 = weight_2311102031
		}
	}

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", minWeight_2311102031)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight_2311102031)
}
