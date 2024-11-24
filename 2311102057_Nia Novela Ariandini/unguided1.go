package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah int
	var weights [1000]float64

	//input banyaknya anak kelinci yang akan ditimbang
	fmt.Print("Masukkan jumlah anak kelinci : ")
	fmt.Scan(&jumlah)

	//input berat anak kelinci
	fmt.Print("Masukkan berat anak kelinci : ")
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&weights[i])
	}

	//menentukan berat terkecil dan terbesar
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for i := 0; i < jumlah; i++ {
		if weights[i] < min {
			min = weights[i] //memperbarui nilai terkecil
		}
		if weights[i] > max {
			max = weights[i] //memperbarui nilai terbesar
		}
	}

	// Output: berat kelinci terkecil dan terbesar
	fmt.Printf("Berat kelinci terkecil : %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar : %.2f\n", max)
}
