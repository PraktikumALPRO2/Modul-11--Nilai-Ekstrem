package main
import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("INPUT JUMLAH WADAH: ")
	fmt.Scan(&x)
	fmt.Print("INPUT JUMLAH IKAN/WADAH: ")
	fmt.Scan(&y)

	berat := make([][]float64, x) 

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan untuk wadah ke-%d:\n", i+1)
		berat[i] = make([]float64, y) 
		for j := 0; j < y; j++ {
			fmt.Scan(&berat[i][j])
		}
	}
	for i := 0; i < x; i++ {
		var jumlah float64
		for j := 0; j < y; j++ {
			jumlah += berat[i][j]
		}
		rata2 := jumlah / float64(y)
		fmt.Printf("Wadah %d: jumlah total = %.2f, rata-rata = %.2f\n", i+1, jumlah, rata2)
	}
}
