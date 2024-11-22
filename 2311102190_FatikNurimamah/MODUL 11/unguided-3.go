package main

import (
    "fmt"
    "math"
)

type ArrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat ArrBalita, JumlahData int, bMin, bMax *float64) {
    // Inisialisasi nilai awal min dan max
    *bMin = math.MaxFloat64
    *bMax = -math.MaxFloat64
    
    // Mencari nilai minimum dan maksimum
    for i := 0; i < JumlahData; i++ {
        if arrBerat[i] < *bMin {
            *bMin = arrBerat[i]
        }
        if arrBerat[i] > *bMax {
            *bMax = arrBerat[i]
        }
    }
}

// Fungsi untuk menghitung rata-rata berat
func rerata(arrBerat ArrBalita, JumlahData int) float64 {
    var total float64 = 0
    
    // Menghitung total berat
    for i := 0; i < JumlahData; i++ {
        total += arrBerat[i]
    }
    
    // Mengembalikan rata-rata
    return total / float64(JumlahData)
}

func main() {
    var balita ArrBalita
    var JumlahBalita int
    var min, max float64
    
    // Input jumlah data
    fmt.Print("Masukkan jumlah data balita: ")
    fmt.Scan(&JumlahBalita)
	fmt.Println()
    
    // Input berat balita
    for i := 0; i < JumlahBalita; i++ {
        fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
        fmt.Scan(&balita[i])	
    }
    
    // Hitung min, max, dan rata-rata
    hitungMinMax(balita, JumlahBalita, &min, &max)
    avg := rerata(balita, JumlahBalita)
    
    // Tampilkan hasil
    fmt.Printf("\nHasil Analisis Data Berat Balita:\n")
    fmt.Printf("Berat Balita Minimum: %.2f kg\n", min)
    fmt.Printf("Berat Balita Maksimum: %.2f kg\n", max)
    fmt.Printf("Berat Rata-rata: %.2f kg\n", avg)
}