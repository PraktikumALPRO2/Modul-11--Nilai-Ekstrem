package main

import ("fmt")

func main() {
	// Membaca jumlah ikan dan kapasitas wadah
	var x, y int
	fmt.Println("Masukkan jumlah ikan dan kapasitas wadah (x y):")
	fmt.Scan(&x, &y)

	// Validasi jumlah ikan dan kapasitas wadah
	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0 dan jumlah ikan maksimal 1000.")
		return
	}

	// Membaca berat ikan
	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan (dipisahkan spasi):")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	// Hitung total berat di setiap wadah
	jumlahWadah := (x + y - 1) / y // Rumus untuk menghitung jumlah wadah (pembulatan ke atas)
	totalBeratPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indexWadah := i / y
		totalBeratPerWadah[indexWadah] += beratIkan[i]
	}

	// Hitung rata-rata berat per wadah
	totalRataRata := 0.0
	for _, totalBerat := range totalBeratPerWadah {
		totalRataRata += totalBerat
	}
	rataRataPerWadah := totalRataRata / float64(jumlahWadah)

	// Tampilkan hasil
	fmt.Println("Total berat di setiap wadah:")
	for _, totalBerat := range totalBeratPerWadah {
		fmt.Printf("%.2f ", totalBerat)
	}
	fmt.Println()
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRataPerWadah)
}