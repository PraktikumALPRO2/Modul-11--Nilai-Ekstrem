// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

func main() {
	// Untuk jumlah ikan (x) dan jumlah ikan per wadah (y)
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah : ")
	fmt.Scan(&x, &y)

	// Mevalidasi jumlah ikan per wadah
	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("Jumlah ikan harus tidak lebih dari 1000 dan jumlah ikan per wadah harus lebih dari 0")
		return
	}

	// array untuk menyimpan berat ikan
	ikan := make([]float64, x) 
	fmt.Println("Masukkan berat ikan :")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i]) 
	}

	
	// Menghitung jumlah wadah yang dibutuhkan
	totWadah := (x + y - 1) / y 
	beratPerWadah := make([]float64, totWadah) 

	// untuk mengisi array beratPerWadah dengan total berat ikan pada setiap wadah
	for i := 0; i < x; i++ {
		idxWadah := i / y 
		beratPerWadah[idxWadah] += ikan[i] 
	}

	// Output total berat ikan pada setiap wadah
	fmt.Println("Total berat ikan pada setiap wadah :")
	for i, berat := range beratPerWadah {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, berat) 
	}

	// untuk menghitung berat rata-rata ikan pada setiap wadah
	var totalBerat float64 
	for _, berat := range beratPerWadah {
		totalBerat += berat 
	}
	rataRataBerat := totalBerat / float64(totWadah) // untuk menghitung rata-rata berat

	// Output rata-rata berat ikan 
	fmt.Printf("Berat rata-rata ikan pada setiap wadah: %.2f kg\n", rataRataBerat)
}
