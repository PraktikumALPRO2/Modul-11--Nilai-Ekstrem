/* Liya Khoirunnisa - 2311102124 */
package main

import (
	"fmt"
)

// Fungsi untuk menghitung total berat ikan di setiap wadah
func hitungTotalBerat_124(wadah []float64, ikan []float64, totalIkan, kapasitasWadah int) {
	for i := 0; i < totalIkan; i += kapasitasWadah {
		var total float64
		for j := i; j < i+kapasitasWadah && j < totalIkan; j++ {
			total += ikan[j]
		}
		wadah[i/kapasitasWadah] = total // Simpan total berat ke wadah
	}
}

// Fungsi untuk menghitung berat rata-rata ikan di setiap wadah
func hitungRataRata(wadah []float64, totalIkan, kapasitasWadah int) []float64 {
	var rataRata []float64
	for i := 0; i < len(wadah); i++ {
		count := 0
		for j := i * kapasitasWadah; j < (i+1)*kapasitasWadah && j < totalIkan; j++ {
			count++
		}
		if count > 0 {
			rataRata = append(rataRata, wadah[i]/float64(count))
		} else {
			rataRata = append(rataRata, 0)
		}
	}
	return rataRata
}

func main() {
	// Deklarasi variabel
	var x, y int

	// Meminta input jumlah ikan dan kapasitas maksimum dalam wadah
	fmt.Print("Masukkan total ikan dan kapasitas maksimum dalam wadah: ")
	fmt.Scan(&x, &y)

	// Cek jika x lebih dari kapasitas array
	if x > 1000 {
		fmt.Println("Jumlah ikan tidak boleh lebih dari 1000.")
		return
	}

	// Membuat array untuk menyimpan berat ikan
	ikan := make([]float64, x)

	// Meminta input berat ikan satu per satu
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		// Menampilkan prompt input dengan nomor ikan
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	// Menghitung jumlah wadah yang diperlukan
	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)

	// Memanggil fungsi untuk menghitung total berat ikan di setiap wadah
	hitungTotalBerat_124(totalBeratWadah, ikan, x, y)

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total := range totalBeratWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, total)
	}

	// Menghitung rata-rata berat ikan di setiap wadah
	rataRata := hitungRataRata(totalBeratWadah, x, y)

	// Menampilkan berat rata-rata ikan di setiap wadah
	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	for i, rata := range rataRata {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, rata)
	}
}
