/* Liya Khoirunnisa - 2311102124 */
package main

import "fmt"

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungBeratKelinci_124(berat []float64, jumlahData int, beratMin, beratMax *float64) {
	*beratMin = berat[0] // Inisialisasi berat minimum
	*beratMax = berat[0] // Inisialisasi berat maksimum

	// Perulangan untuk mencari berat kelinci yang minimum dan maksimum
	for i := 1; i < jumlahData; i++ {
		if berat[i] < *beratMin { // Jika berat kelinci lebih kecil dari berat minimum
			*beratMin = berat[i] // Update berat minimum
		}
		if berat[i] > *beratMax { // Jika berat kelinci lebih besar dari berat maksimum
			*beratMax = berat[i] // Update berat maksimum
		}
	}
}

func main() {
	// Deklarasi variabel
	var jumlahKelinci int

	// Menggunakan slice untuk data berat kelinci
	var dataKelinci [1000]float64 // Array dengan kapasitas 1000

	// Deklarasi variabel
	var beratMin, beratMax float64

	// Meminta input jumlah data berat kelinci
	fmt.Print("Masukkan jumlah data berat kelinci: ")
	fmt.Scan(&jumlahKelinci)

	// Validasi jumlah data kelinci
	if jumlahKelinci < 1 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah data harus antara 1 dan 1000.")
		return
	}

	// Input berat kelinci
	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&dataKelinci[i])
	}

	// Menampilkan hasil perhitungan berat minimum dan maksimum
	hitungBeratKelinci_124(dataKelinci[:], jumlahKelinci, &beratMin, &beratMax)

	// Menampilkan hasil
	fmt.Print("\nHitung Berat Kelinci\n")
	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", beratMin)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", beratMax)
}
