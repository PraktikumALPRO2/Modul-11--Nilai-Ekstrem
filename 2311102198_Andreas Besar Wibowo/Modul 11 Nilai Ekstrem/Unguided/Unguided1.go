// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Tipe data array untuk menyimpan berat kelinci
type Kelinci struct {
	Berat float64
}

func main() {
	// Untuk Inputan jumlah anak kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	// Mevalidasi jumlah anak kelinci
	if N < 1 || N > 1000 {
		fmt.Println("Jumlah anak kelinci tidak lebih dari 1000.")
		return
	}

	// untuk Array Kelinci
	daftarKelinci := make([]Kelinci, N)

	// Untuk Inputan berat anak-anak kelinci
	fmt.Println("Masukkan berat anak-anak kelinci:")
	for i := 0; i < N; i++ {
		var berat float64
		fmt.Printf("Berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat)
		// Memasukkan berat kelinci ke Array Kelinci
		daftarKelinci[i] = Kelinci{Berat: berat}
	}

	minKelinci := daftarKelinci[0] // untuk nilai min dengan berat kelinci pertama
	maxKelinci := daftarKelinci[0] // untuk nilai max dengan berat kelinci pertama

	// untuk mencari berat min dan max kelinci
	for _, kelinci := range daftarKelinci {
		if kelinci.Berat < minKelinci.Berat {
			minKelinci = kelinci
		}
		if kelinci.Berat > maxKelinci.Berat {
			maxKelinci = kelinci
		}
	}

	// Untuk Output Akhir
	fmt.Printf("Berat terkecil: %.2f kg\n", minKelinci.Berat)
	fmt.Printf("Berat terbesar: %.2f kg\n", maxKelinci.Berat)
}
