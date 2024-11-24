package main

import "fmt"

type kelinci [1000]float64

func isiArray(a *kelinci, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&a[i])
	}
}

func cariTerkecil(a kelinci, n int) float64 {
	terkecil := a[0]
	for i := 1; i < n; i++ {
		if terkecil > a[i] {
			terkecil = a[i]
		}
	}
	return terkecil
}

func cariTerbesar(a kelinci, n int) float64 {
	terbesar := a[0]
	for i := 1; i < n; i++ {
		if terbesar < a[i] {
			terbesar = a[i]
		}
	}
	return terbesar
}

func main() {
	var berat kelinci
	var n int

	fmt.Print("Banyak kelinci ditimbang: ")
	fmt.Scan(&n)

	isiArray(&berat, n)

	fmt.Printf("Kelinci terkecil beratnya sebesar %.2f sedangkan kelinci terbesar beratnya sebesar %.2f.\n",
		cariTerkecil(berat, n), cariTerbesar(berat, n))
}