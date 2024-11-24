package main

import (
	"fmt"
)

func main() {
	var jumlahIkan, kapasitasWadah int

	// Input jumlah ikan yang dijual dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan yang dijual: ")
	fmt.Scan(&jumlahIkan)
	fmt.Print("Masukkan kapasitas wadah (jumlah ikan per wadah): ")
	fmt.Scan(&kapasitasWadah)

	// Validasi input
	if jumlahIkan <= 0 || kapasitasWadah <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari nol.")
		return
	}

	// Input berat masing-masing ikan
	beratIkan := make([]float64, jumlahIkan)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Pengelompokan ikan ke dalam wadah
	var totalBeratPerWadah []float64
	for i := 0; i < jumlahIkan; i += kapasitasWadah {
		batasAkhir := i + kapasitasWadah
		if batasAkhir > jumlahIkan {
			batasAkhir = jumlahIkan
		}

		// Hitung total berat ikan dalam satu wadah
		var total float64
		for j := i; j < batasAkhir; j++ {
			total += beratIkan[j]
		}
		totalBeratPerWadah = append(totalBeratPerWadah, total)
	}

	// Menampilkan hasil total berat dan rata-rata berat setiap wadah
	fmt.Println("\nHasil pengelompokan ikan:")
	for i, total := range totalBeratPerWadah {
		rataRata := total / float64(kapasitasWadah)
		if i == len(totalBeratPerWadah)-1 && jumlahIkan%kapasitasWadah != 0 {
			rataRata = total / float64(jumlahIkan%kapasitasWadah) // Penyesuaian untuk wadah terakhir
		}
		fmt.Printf("Wadah ke-%d: Total berat = %.2f kg, Rata-rata = %.2f kg\n", i+1, total, rataRata)
	}
}
