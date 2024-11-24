package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x > 1000 {
		fmt.Println("Jumlah ikan tidak boleh lebih dari 1000")
		return
	}

	ikan := make([]float64, x)
	fmt.Println("Masukkan berat setiap ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	wadah := (x + y - 1) / y 
	totalBeratPerWadah := make([]float64, wadah)
	for i := 0; i < x; i++ {
		totalBeratPerWadah[i/y] += ikan[i]
	}

	fmt.Println("Total berat di setiap wadah:")
	for i, total := range totalBeratPerWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	var totalBerat float64
	for _, total := range totalBeratPerWadah {
		totalBerat += total
	}
	rataRata := totalBerat / float64(wadah)
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRata)
}
