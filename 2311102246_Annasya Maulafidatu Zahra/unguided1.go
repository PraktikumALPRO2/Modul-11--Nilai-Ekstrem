package main

import "fmt"

func main() {
	var jumlahKelinci int
	var berat [1000]float64
	var beratTerkecil, beratTerbesar float64

	fmt.Print("Masukkan Jumlah Kelinci: ")
	fmt.Scan(&jumlahKelinci)

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	beratTerkecil = berat[0]
	beratTerbesar = berat[0]
	for i := 1; i < jumlahKelinci; i++ {
		if beratTerbesar < berat[i] {
			beratTerbesar = berat[i]
		}

		if beratTerkecil > berat[i] {
			beratTerkecil = berat[i]
		}
	}

	fmt.Println("Berat kelinci terbesar adalah: ", beratTerbesar, "kg")
	fmt.Println("Berat kelinci terkecil adalah: ", beratTerkecil, "kg")
}
