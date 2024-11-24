package main

import "fmt"

type arrBeratAnak [100]float64

func cariMinMax(beratAnak arrBeratAnak, beratTerkecil, beratTerbesar *float64) {
	*beratTerkecil = beratAnak[0]
	*beratTerbesar = beratAnak[0]

	for i := 1; i < len(beratAnak); i++ {
		if beratAnak[i] != 0 {
			if beratAnak[i] < *beratTerkecil {
				*beratTerkecil = beratAnak[i]
			}
			if beratAnak[i] > *beratTerbesar {
				*beratTerbesar = beratAnak[i]
			}
		}
	}
}

func hitungRataRata(beratAnak arrBeratAnak, jumlahAnak int) float64 {
	var totalBerat float64
	for i := 0; i < jumlahAnak; i++ {
		totalBerat += beratAnak[i]
	}
	return totalBerat / float64(jumlahAnak)
}

func main() {
	var beratAnak arrBeratAnak
	var jumlahAnak int
	var beratTerkecil, beratTerbesar float64

	fmt.Print("Masukkan jumlah anak : ")
	fmt.Scan(&jumlahAnak)

	for i := 0; i < jumlahAnak; i++ {
		fmt.Printf("Masukkan berat anak ke-%d (kg) : ", i+1)
		fmt.Scan(&beratAnak[i])
	}

	cariMinMax(beratAnak, &beratTerkecil, &beratTerbesar)

	rataRataBerat := hitungRataRata(beratAnak, jumlahAnak)

	fmt.Printf("\n= Hasil Analisis Data Berat Anak =\n")
	fmt.Printf("Berat Terkecil : %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat Terbesar : %.2f kg\n", beratTerbesar)
	fmt.Printf("Berat Rata-rata : %.2f kg\n", rataRataBerat)
}
