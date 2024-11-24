package main

import "fmt"

type ikan [1000]float64

func dataikan(b *ikan, x int) {
	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&b[i])
	}
}

func totalBeratDanRataRata(b *ikan, x int, y int) ([]float64, []float64) {
	var totalBerat []float64
	var rataRata []float64

	totalWadah := (x + y - 1) / y

	for wadah := 0; wadah < totalWadah; wadah++ {
		start := wadah * y
		end := start + y
		if end > x {
			end = x
		}

		var total float64
		for i := start; i < end; i++ {
			total += b[i]
		}
		totalBerat = append(totalBerat, total)

		jumlahIkan := end - start
		rataRata = append(rataRata, total/float64(jumlahIkan))
	}

	return totalBerat, rataRata
}

func main() {
	var ikanData ikan
	var x, y int

	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan kapasitas per wadah: ")
	fmt.Scan(&y)

	fmt.Println("Silahkan masukkan berat ikan!")
	dataikan(&ikanData, x)

	totalBerat, rataRata := totalBeratDanRataRata(&ikanData, x, y)

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, berat := range totalBerat {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()

	fmt.Println("Rata-rata berat ikan di setiap wadah:")
	for _, rata := range rataRata {
		fmt.Printf("%.2f ", rata)
	}
	fmt.Println()
}