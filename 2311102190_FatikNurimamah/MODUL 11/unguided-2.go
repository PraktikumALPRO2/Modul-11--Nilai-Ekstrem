package main

import "fmt"

func main() {
    var jumlahIkan, ikanPerWadah int
    var beratIkanArray [1000]float64
    var totalBeratPerWadah []float64 // untuk menyimpan total berat per wadah
    
    // Input banyak ikan (jumlahIkan) dan banyak ikan per wadah (ikanPerWadah)
    fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah: ")
    fmt.Scan(&jumlahIkan, &ikanPerWadah)
    
    // Input berat masing-masing ikan
    fmt.Println("Masukkan berat masing-masing ikan:")
    for i := 0; i < jumlahIkan; i++ {
        fmt.Scan(&beratIkanArray[i])
    }
    
    // Hitung total berat per wadah
    jumlahWadah := (jumlahIkan + ikanPerWadah - 1) / ikanPerWadah // rumus pembulatan ke atas
    var indeksIkan int = 0
    
    for i := 0; i < jumlahWadah; i++ {
        var totalBeratWadah float64 = 0
        var jumlahIkanDalamWadah int = 0
        
        // Hitung total berat untuk wadah
        for j := 0; j < ikanPerWadah && indeksIkan < jumlahIkan; j++ {
            totalBeratWadah += beratIkanArray[indeksIkan]
            indeksIkan++
            jumlahIkanDalamWadah++
        }
        totalBeratPerWadah = append(totalBeratPerWadah, totalBeratWadah)
    }
    
    // Output total berat per wadah
    fmt.Println("\nTotal berat ikan di setiap wadah:")
    for i, totalBerat := range totalBeratPerWadah {
        fmt.Printf("Wadah %d: %.2f\n", i+1, totalBerat)
    }
    
    // Hitung rata-rata berat ikan di semua wadah
    var totalBeratSemuaWadah float64 = 0
    for _, totalBerat := range totalBeratPerWadah {
        totalBeratSemuaWadah += totalBerat
    }
    rataRataBeratPerWadah := totalBeratSemuaWadah / float64(len(totalBeratPerWadah))
    
    fmt.Printf("\nRata-rata berat ikan per wadah: %.2f\n", rataRataBeratPerWadah)
}
