package main

import "fmt"

func main() {
	var x_158, y_158 int

	fmt.Print("Masukkan total ikan dan kapasitas maksimum dalam wadah: ")
	fmt.Scanln(&x_158, &y_158)

	if x_158 > 0 && x_158 <= 1000 {
		berat := make([]float64, x_158)
		jumlahIkan := make([]int, x_158)

		fmt.Println("Masukkan berat ikan:")
		for i := 0; i < x_158; i++ {
			fmt.Printf("Berat ikan ke-%d: ", i+1)
			fmt.Scanln(&berat[i])
			jumlahIkan[i] = y_158
		}

		totalBerat := make([]float64, x_158)
		for i := 0; i < x_158; i++ {
			totalBerat[i] = berat[i] * float64(jumlahIkan[i])
		}

		fmt.Println("\nTotal berat ikan di setiap wadah:")
		for i := 0; i < x_158; i++ {
			fmt.Printf("Wadah %d: %.2f kg\n", i+1, totalBerat[i])
		}

		fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
		for i := 0; i < x_158; i++ {
			rataRataPerWadah := berat[i] 
			fmt.Printf("Wadah %d: %.2f kg\n", i+1, rataRataPerWadah)
		}

	} else {
		fmt.Println("Jumlah wadah harus antara 1 hingga 1000.")
	}
}
