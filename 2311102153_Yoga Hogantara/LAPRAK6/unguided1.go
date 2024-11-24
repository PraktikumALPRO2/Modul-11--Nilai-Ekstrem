package main
import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("INPUT JUMALH ANAK KELINCI: ")
	fmt.Scan(&n)
	if n > 1000{
		fmt.Println("MAKS 1000")
		return
	}
	berat:=make([]float64,n)

	fmt.Println("INPUT BERAT KELINCI:")
	for i:=0;i<n;i++{
		fmt.Scan(&berat[i])
	}
	max,min:=berat[0],berat[0]

	for i := 0; i < n; i++ {
		if berat[i]<min{
			min=berat[i]
		}
		if berat[i]>max{
			max=berat[i]
		}
	}
	fmt.Println("BERAT KELINCI TERBESAR :",max)
	fmt.Println("BERAT KELINCI TERKECIL :",min)
}