package main

import "fmt"

func main() {
	var n int
	var berat []float64

	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scan(&n)

	//input berat anak kelinci
	berat = make([]float64, n)
	fmt.Print("Masukkan berat masisng - masing anak kelinci: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	//inisialisasi nilai extreme
	beratMinim := berat[0]
	beratMaks := berat[0]

	//pencarian nilai extreme
	for _, berat := range berat {
		if berat < beratMinim {
			beratMinim = berat
		}
		if berat > beratMaks {
			beratMaks = berat
		}
	}

	//output
	fmt.Printf("Berat terkecil: %.1f\n", beratMinim)
	fmt.Printf("Berat terbesar: %.1f\n", beratMaks)
}
