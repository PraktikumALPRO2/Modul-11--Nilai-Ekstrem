//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	beratKelinci := make([]float64, N)
	fmt.Printf("Masukkan berat %d anak kelinci:\n", N)
	for i := 0; i < N; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	for _, berat := range beratKelinci {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}
