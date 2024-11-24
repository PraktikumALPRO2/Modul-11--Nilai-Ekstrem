package main

import "fmt"

func main() {
	var n_158 int
	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scanln(&n_158)

	if n_158 > 0 && n_158 <= 1000 {
		berat := make([]float64, n_158)

		fmt.Println("Masukkan berat anak kelinci:")
		for i := 0; i < n_158; i++ {
			fmt.Scan(&berat[i])
		}

		terkecil := berat[0]
		terbesar := berat[0]

		for i := 1; i < n_158; i++ {
			if berat[i] < terkecil {
				terkecil = berat[i]
			}
			if berat[i] > terbesar {
				terbesar = berat[i]
			}
		}

		fmt.Printf("\nBerat terkecil: %.2f\n", terkecil)
		fmt.Printf("Berat terbesar: %.2f\n", terbesar)
	} else {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
	}
}
