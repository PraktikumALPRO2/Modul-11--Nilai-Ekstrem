//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus valid (x <= 1000, y > 0).")
		return
	}

	fmt.Printf("Masukkan berat %d ikan:\n", x)
	beratIkan := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalBerat float64
	wadah := []float64{}
	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]
		if (i+1)%y == 0 || i == x-1 {
			wadah = append(wadah, totalBerat)
			totalBerat = 0
		}
	}

	var totalWadah float64
	for _, berat := range wadah {
		totalWadah += berat
	}
	rataRata := totalWadah / float64(len(wadah))

	fmt.Println("Total berat di setiap wadah:")
	for _, berat := range wadah {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}
