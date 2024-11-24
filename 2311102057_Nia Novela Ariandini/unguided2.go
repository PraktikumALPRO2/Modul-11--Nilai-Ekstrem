package main

import "fmt"

func main() {
    var x, y int

    //input jumlah ikan dan jumlah ikan per wadah
    fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah : ")
    fmt.Scan(&x, &y)

    //menyimpan berat ikan dalam array
    var beratIkan [1000]float64
    fmt.Print("Masukkan berat ikan : ")
    for i := 0; i < x; i++ {
        fmt.Scan(&beratIkan[i])
    }

    //menghitung total berat ikan di setiap wadah
    var totalBerat []float64
    var beratWadah float64

    for i := 0; i < x; i++ {
        beratWadah += beratIkan[i] //tambahkan berat ikan ke wadah saat ini

        //jika sudah cukup banyak ikan dalam wadah (y ikan), simpan total berat wadah
        if (i+1)%y == 0 || i == x-1 {
            totalBerat = append(totalBerat, beratWadah)
            beratWadah = 0 //reset berat wadah setelah memasukkan ikan
        }
    }

    //ouput : total berat ikan di setiap wadah
    fmt.Print("Total berat ikan di setiap wadah : ")
    for _, berat := range totalBerat {
        fmt.Printf("%.2f ", berat)
    }
    fmt.Println()

    //menghitung rata-rata berat ikan per wadah
    rataRata := 0.0
    for _, berat := range totalBerat {
        rataRata += berat
    }

    rataRata /= float64(len(totalBerat))

    //output : rata-rata berat ikan per wadah
    fmt.Printf("Rata-rata berat ikan di setiap wadah : %.2f\n", rataRata)
}
