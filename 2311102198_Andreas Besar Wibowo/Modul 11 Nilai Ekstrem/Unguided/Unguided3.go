// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Tipe data array untuk menyimpan berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat min dan max
func hitungMinMax(arrBerat arrBalita, n int) (minBerat, maxBerat float64) {
	if n == 0 {
		return 0, 0 
	}

	// untuk nilai minimum dan maksimum
	minBerat, maxBerat = arrBerat[0], arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < minBerat {
			minBerat = arrBerat[i]
		}
		if arrBerat[i] > maxBerat {
			maxBerat = arrBerat[i]
		}
	}
	return
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arrBerat arrBalita, n int) float64 {
	if n == 0 {
		return 0
	}

	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var dataBerat arrBalita

	for i := 0; i < len(dataBerat); i++ {
		dataBerat[i] = 0
	}

	// Input jumlah balita yang ingin di data
	fmt.Print("Banyak data berat balita : ")
	fmt.Scan(&n)

	// mevalidasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data tidak lebih dari 100.")
		return
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke - %d : ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	// untuk menghitung minimum, maksimum, dan rata-rata berat
	minBerat, maxBerat := hitungMinMax(dataBerat, n)
	rataRata := hitungRerata(dataBerat, n)

	// Output akhir
	fmt.Printf("\nBerat Balita Minimum : %.2f kg\n", minBerat)
	fmt.Printf("Berat Balita Maksimal : %.2f kg\n", maxBerat)
	fmt.Printf("Berat Balita Rata-Rata : %.2f kg\n", rataRata)
}
