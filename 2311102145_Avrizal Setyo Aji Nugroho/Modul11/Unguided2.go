package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Input (X) jumlah ikan  dan (y) kapasitas ikan per wadah : ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Input tidak valid!")
		return
	}

	fmt.Println("Input berat ikan :")
	beratIkan_145 := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan_145[i])
	}

	var totalWadah_145 []float64
	var totalBerat_145 float64
	for i := 0; i < x; i++ {
		if i%y == 0 {
			totalWadah_145 = append(totalWadah_145, 0)
		}
		totalWadah_145[len(totalWadah_145)-1] += beratIkan_145[i]
		totalBerat_145 += beratIkan_145[i]
	}

	rataRata := totalBerat_145 / float64(len(totalWadah_145))

	fmt.Println("Total berat per wadah:")
	for i, berat := range totalWadah_145 {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRata)
}
