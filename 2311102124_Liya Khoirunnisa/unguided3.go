/* Liya Khoirunnisa - 2311102124 */
package main

import "fmt"

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungBeratBalita_124(berat []float64, jumlahData int, beratMin, beratMax *float64) {
	*beratMin = berat[0] // Inisialisasi berat minimum
	*beratMax = berat[0] // Inisialisasi berat maksimum

	// Perulangan untuk mencari berat balita yang minimum dan maksimum
	for i := 1; i < jumlahData; i++ {
		if berat[i] < *beratMin { // Jika berat balita lebih kecil dari berat minimum
			*beratMin = berat[i] // Update berat minimum
		}
		if berat[i] > *beratMax { // Jika berat balita lebih besar dari berat maksimum
			*beratMax = berat[i] // Update berat maksimum
		}
	}
	// Menampilkan hasil berat minimum dan maksimum
	fmt.Printf("Berat balita minimum: %.2f kg\n", *beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *beratMax)
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRataRata(berat []float64, jumlahData int) float64 {
	totalBerat := 0.0 // Variabel untuk menyimpan total berat balita

	// Perulangan untuk menghitung total berat dari semua balita
	for i := 0; i < jumlahData; i++ {
		totalBerat += berat[i]
	}

	// Menghitung rata-rata berat balita
	rataRata := totalBerat / float64(jumlahData)

	// Menampilkan hasil rata-rata berat
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
	return rataRata // Mengembalikan rata-rata berat
}

func main() {
	// Deklarasi variabel
	var jumlahBalita int

	// Menggunakan slice untuk data berat balita
	var dataBalita []float64

	// Deklarasi variabel
	var beratMin, beratMax float64

	// Meminta input jumlah data berat balita
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahBalita)

	// Validasi jumlah data balita
	if jumlahBalita > 100 {
		fmt.Println("Jumlah data tidak boleh lebih dari 100.")
		return
	}

	// slice dengan panjang sesuai jumlah data
	dataBalita = make([]float64, jumlahBalita)

	// Input berat balita
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	// Menampilkan hasil perhitungan berat minimum, maksimum, dan rata-rata
	fmt.Println("\nHasil Perhitungan Berat Balita: ")

	// Panggil fungsi untuk menghitung berat minimum dan maksimum
	hitungBeratBalita_124(dataBalita, jumlahBalita, &beratMin, &beratMax)

	// Panggil fungsi untuk menghitung dan menampilkan rata-rata berat
	hitungRataRata(dataBalita, jumlahBalita)
}
