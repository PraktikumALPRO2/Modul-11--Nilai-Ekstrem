<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 11</strong></h2>
<h2 align="center"><strong> NILAI EKSTRIM </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Unguided](#unguided)
3. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori
**1. Pencarian Nilai Ekstrim pada Array Bertipe Data Dasar**

Pencarian nilai ekstrim pada array bertipe data dasar adalah proses untuk menemukan dan menentukkan nilai maksimum (nilai tertinggi) dan nilai minimum (nilai terendah) dalam sebuah array yang berisi elemen-elemen dari tipe data dasar, seperti integer, float, atau string.

**Langkah-langkahnya adalah:**  

1. Inisialisasi: Tetapkan dua variabel untuk menyimpan nilai maksimum dan minimum. Biasanya, keduanya diinisialisasi menggunakan elemen pertama dari array.  
2. Iterasi: Gunakan loop untuk menelusuri setiap elemen dalam array.  
3. Perbandingan: Selama iterasi, bandingkan setiap elemen dengan nilai maksimum dan minimum yang tersimpan. Jika elemen lebih besar dari nilai maksimum, perbarui nilai maksimum. Jika elemen lebih kecil dari nilai minimum, perbarui nilai minimum.  
4. Output: Setelah iterasi selesai, tampilkan nilai maksimum dan minimum yang telah ditemukan.

```go
numbers := []int{34, 12, 24, 56, 5, 78, 90, 1, 23}
```

**2. Pencarian Nilai Ekstrim pada Array Bertipe Data Terstruktur**

Pencarian nilai ekstrem dalam array bertipe data terstruktur adalah proses untuk menemukan dan menentukkan nilai maksimum dan nilai minimum berdasarkan atribut tertentu dari struktur data, seperti objek atau struct.

**Langkah-langkahnya adalah:**  

1. Definisi Struktur: Mulailah dengan membuat sebuah struktur data yang mencakup atribut-atribut yang akan dianalisis. Sebagai contoh, gunakan struktur `Person` dengan atribut seperti `Name` dan `Age`.  
2. Inisialisasi: Sama seperti pada array bertipe data dasar, tetapkan variabel awal untuk menyimpan nilai maksimum dan minimum berdasarkan atribut tertentu yang relevan.  
3. Iterasi: Gunakan loop untuk memproses setiap elemen dalam array struktur tersebut.  
4. Perbandingan: Lakukan perbandingan antara atribut yang dipilih (contohnya, `Age`) dengan nilai maksimum dan minimum yang sudah disimpan. Jika ditemukan nilai yang lebih besar atau lebih kecil, perbarui variabel maksimum atau minimum.  
5. Output: Setelah iterasi selesai, tampilkan hasil akhir dari analisis nilai maksimum dan minimum.

```go
type Person struct {
	Name string
	Age  int
}
people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}, {"David", 20}}
```


  

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan rill berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
### Source Code :
```go
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



```
### Output:
![image](https://github.com/user-attachments/assets/b340e5ca-3bc5-49c8-be71-00c0c5a04783)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/0de983b5-c081-40fd-9efb-0d60284234c8)

### Deskripsi Program : 
Program tersebut adalah program yang digunakan untuk menentukan berat terkecil dan terbesar dari beberapa kelinci. Data berat setiap kelinci disimpan dalam array yang mampu menampung hingga 1000 elemen. Pengguna diminta untuk memasukkan jumlah kelinci, lalu menginput berat masing-masing kelinci satu per satu. Program kemudian memproses data tersebut menggunakan perulangan untuk mencari berat terkecil dan terbesar dengan membandingkan setiap elemen dalam array. Hasil akhirnya ditampilkan dalam bentuk berat kelinci terkecil dan terbesar sesuai data yang dimasukkan.

### 2.  Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program Ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (Jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

### Source Code :
```go
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

```
### Output:
![image](https://github.com/user-attachments/assets/8191b50f-dbef-42cd-9a29-5f371667b302)

### Full code Screenshot:
![code](https://github.com/user-attachments/assets/193ed06e-5b6b-4b6e-90f2-b5d7189335df)

### Deskripsi Program : 
Program tersebut adalah program yang berfungsi untuk menghitung total berat ikan di setiap wadah dan rata-rata berat ikan per wadah. Pengguna diminta untuk memasukkan jumlah ikan secara keseluruhan dan kapasitas ikan yang dapat dimasukkan ke dalam satu wadah. Berat masing-masing ikan dimasukkan ke dalam array, kemudian program akan menghitung total berat ikan pada tiap wadah dengan cara membagi ikan secara berurutan sesuai kapasitas yang ditentukan. Hasil dari program ini berupa daftar total berat ikan per wadah serta rata-rata berat ikan di seluruh wadah. Perhitungan dilakukan menggunakan perulangan untuk memastikan pembagian ikan dan penghitungan berat dilakukan secara sistematis.
 
### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

Buatlah program dengan spesifikasi subprogram sebagai berikut:

![Screenshot 2024-11-22 213235](https://github.com/user-attachments/assets/54c18c89-d4ac-4afc-beba-cc7e0c687c18)
![Screenshot 2024-11-22 213257](https://github.com/user-attachments/assets/e23f1824-9d87-4275-9250-61efd423c742)


### Source Code :
```go
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
		
```
### Output:
![image](https://github.com/user-attachments/assets/00b3a2cf-1429-4183-812c-f6b804878cd8)

### Full code Screenshot:
![image](https://github.com/user-attachments/assets/c378cc02-465f-44aa-92be-b567884804ef)

### Deskripsi Program : 
Program tersebut adalah program yang dirancang untuk menganalisis data berat badan balita. Pengguna diminta untuk memasukkan jumlah balita dan berat masing-masing balita dalam satuan kilogram. Program akan menghitung berat badan terendah dan tertinggi menggunakan fungsi `hitungMinMax`, serta menghitung rata-rata berat badan dengan fungsi `rerata`. Proses analisis dilakukan dengan cara menjumlahkan data berat badan serta membandingkan nilai-nilainya melalui perulangan. Hasil akhir yang ditampilkan ke layar berupa informasi berat balita minimum, berat balita maksimum, dan berat rata-rata nya.

## Daftar Pustaka
[1] Agung, N. (n.d.). Array. Retrieved from Dasar Pemrograman Golang: https://dasarpemrogramangolang.novalagung.com/A-array.html
