package main

import (
	"fmt"
)

func main() {

	var beratIkan []float64

	var jumlahIkan int
	fmt.Print("Masukkan jumlah ikan (x): ")
	fmt.Scan(&jumlahIkan)

	var kapasitasWadah int
	fmt.Print("Masukkan kapasitas wadah (y): ")
	fmt.Scan(&kapasitasWadah)

	if jumlahIkan <= 0 || kapasitasWadah <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
		return
	}

	beratIkan = make([]float64, jumlahIkan)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (jumlahIkan + kapasitasWadah - 1) / kapasitasWadah

	beratWadah := make([]float64, jumlahWadah)
	for i := 0; i < jumlahIkan; i++ {
		wadahKe := i / kapasitasWadah
		beratWadah[wadahKe] += beratIkan[i]
	}

	fmt.Println("Berat total di setiap wadah:")
	for i, berat := range beratWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, berat)
	}

	var totalBerat float64
	for _, berat := range beratWadah {
		totalBerat += berat
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", totalBerat/float64(jumlahWadah))
}
