package main

import (
	"fmt"
)

func validasi(jumlahIkan, kapasitasWadah int) bool {
	return jumlahIkan > 0 && kapasitasWadah > 0
}

func hitungTarif(jumlahIkan, kapasitasWadah int, beratIkan []float64) ([]float64, float64) {
	var totalBeratPerWadah []float64
	var totalKeseluruhan float64

	for i := 0; i < jumlahIkan; i += kapasitasWadah {
		end := i + kapasitasWadah
		if end > jumlahIkan {
			end = jumlahIkan
		}

		var totalBerat float64
		for j := i; j < end; j++ {
			totalBerat += beratIkan[j]
		}
		totalBeratPerWadah = append(totalBeratPerWadah, totalBerat)
		totalKeseluruhan += totalBerat
	}

	rataRataPerWadah := totalKeseluruhan / float64(len(totalBeratPerWadah))
	return totalBeratPerWadah, rataRataPerWadah
}

func main() {
	var jumlahIkan, kapasitasWadah int

	fmt.Print("Masukkan jumlah ikan yang akan dijual: ")
	fmt.Scan(&jumlahIkan)
	fmt.Print("Masukkan kapasitas ikan per wadah: ")
	fmt.Scan(&kapasitasWadah)

	if !validasi(jumlahIkan, kapasitasWadah) {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
		return
	}

	beratIkan := make([]float64, jumlahIkan)
	fmt.Println("Masukkan berat masing-masing ikan (kg):")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	totalBeratPerWadah, rataRataPerWadah := hitungTarif(jumlahIkan, kapasitasWadah, beratIkan)

	fmt.Println("\nTotal berat ikan di setiap wadah (kg):")
	for i, totalBerat := range totalBeratPerWadah {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, totalBerat)
	}
	fmt.Printf("\nRata-rata berat ikan per wadah: %.2f kg\n", rataRataPerWadah)
}
