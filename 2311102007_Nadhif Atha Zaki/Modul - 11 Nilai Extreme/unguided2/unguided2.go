package main

import "fmt"

func main() {
	// Deklarasi variabel
	var x, y int
	var ikan [1000]float64

	// Input jumlah ikan dan kapasitas wadah
	fmt.Print("Masukkan banyaknya ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	// Input berat ikan
	fmt.Printf("Masukkan berat %d ikan: ", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	// Proses penghitungan
	var totalBerat []float64
	for i := 0; i < x; i += y {
		jumlah := 0.0
		for j := i; j < i+y && j < x; j++ {
			jumlah += ikan[j]
		}
		totalBerat = append(totalBerat, jumlah)
	}

	// Hitung rata-rata berat per wadah
	jumlahTotal := 0.0
	for _, berat := range totalBerat {
		jumlahTotal += berat
	}
	rataRata := jumlahTotal / float64(len(totalBerat))

	// Output
	fmt.Println("Total berat di setiap wadah:")
	for _, berat := range totalBerat {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()

	fmt.Printf("Rata-rata berat di setiap wadah: %.2f\n", rataRata)
}
