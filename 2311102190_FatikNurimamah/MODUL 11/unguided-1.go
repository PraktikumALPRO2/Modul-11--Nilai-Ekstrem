package main

import "fmt"

func main() {
    // Deklarasi array untuk menyimpan berat kelinci dengan kapasitas 1000
    var beratKelinci [1000]float64
    var jumlahKelinci int // untuk menyimpan jumlah kelinci yang akan dimasukkan
    
    // Input jumlah kelinci
    fmt.Print("Masukkan jumlah kelinci: ")
    fmt.Scan(&jumlahKelinci)
    fmt.Println()
    
    // Input berat masing-masing kelinci
    for i := 0; i < jumlahKelinci; i++ {
        fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
        fmt.Scan(&beratKelinci[i])
    }
    
    // Mencari berat terkecil dan terbesar
    var beratTerkecil float64 = beratKelinci[0]
    var beratTerbesar float64 = beratKelinci[0]
    
    for i := 1; i < jumlahKelinci; i++ {
        // Mencari berat terkecil
        if beratTerkecil > beratKelinci[i] {
            beratTerkecil = beratKelinci[i]
        }
        // Mencari berat terbesar 
        if beratTerbesar < beratKelinci[i] {
            beratTerbesar = beratKelinci[i]
        }
    }
    
    // Output hasil
    fmt.Printf("\nBerat kelinci terkecil: %g\n", beratTerkecil)
    fmt.Printf("Berat kelinci terbesar: %g\n", beratTerbesar)
}
