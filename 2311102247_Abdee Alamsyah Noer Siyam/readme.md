<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL XI</strong>
  <br>
  <strong>PENCARIAN NILAI EKSTRIN PADA HIMPUNAN DATA</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
  <br>
  S1IF1105
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

<strong><h3>NILAI EKSTRIM</h3></strong>

<strong><h4>PENGERTIAN NILAI EKSTRIM</h4></strong>

Nilai ekstrim adalah nilai maksimum dan minimum dalam sekumpulan data. Dalam bahasa pemrograman, menemukan nilai ekstrim penting untuk analisis data, statistik, dan pengambilan keputusan. Di Go (Golang), kita dapat menghitung nilai ekstrim dari array atau slice dengan menggunakan perulangan dan kondisi sederhana.

<strong><h4>DEKLARASI NILAI EKSTRIM</h4></strong>

Untuk menghitung nilai ekstrim, kita biasanya menggunakan array atau slice. Berikut adalah contoh deklarasi dan inisialisasi slice di Go:

```go
package main

import "fmt"

func main() {
    // Inisialisasi slice dengan beberapa nilai
    numbers := []int{3, 5, 1, 8, 2}
    fmt.Println("Slice:", numbers)
}
```

<strong><h4>MENCARI NILAI MAKSIMUM DAN MINIMUM</h4></strong>
Untuk menemukan nilai maksimum dan minimum dari sebuah slice, kita dapat menggunakan perulangan untuk membandingkan setiap elemen. Berikut adalah contoh fungsi untuk mencari nilai maksimum dan minimum:

```go
package main

import "fmt"

// Fungsi untuk mencari nilai maksimum dan minimum
func findMinMax(numbers []int) (int, int) {
    min := numbers[0]
    max := numbers[0]
    
    for _, number := range numbers {
        if number < min {
            min = number
        }
        if number > max {
            max = number
        }
    }
    return min, max
}

func main() {
    numbers := []int{3, 5, 1, 8, 2}
    min, max := findMinMax(numbers)
    
    fmt.Println("Nilai Minimum:", min) // Output: Nilai Minimum: 1
    fmt.Println("Nilai Maksimum:", max) // Output: Nilai Maksimum: 8
}
```

<strong><h4>CONTOH PENGGUNAAN</h4></strong>
Berikut adalah contoh penggunaan fungsi findMinMax dalam program Go:

```go
package main

import "fmt"

func findMinMax(numbers []int) (int, int) {
    min := numbers[0]
    max := numbers[0]
    
    for _, number := range numbers {
        if number < min {
            min = number
        }
        if number > max {
            max = number
        }
    }
    return min, max
}

func main() {
    // Inisialisasi slice dengan beberapa nilai
    numbers := []int{3, 5, 1, 8, 2}
    
    // Mencari nilai maksimum dan minimum
    min, max := findMinMax(numbers)
    
    // Menampilkan hasil
    fmt.Println("Nilai Minimum:", min) // Output: Nilai Minimum: 1
    fmt.Println("Nilai Maksimum:", max) // Output: Nilai Maksimum: 8
}
```

## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

**Masukan** terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat 𝑁 yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya 𝑁 bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

**Keluaran** terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### Source Code

```go
package main

import "fmt"

func main() {
    var banyak int
    var beratKelinci [1000]float64
    var terkecil, terbesar float64

    fmt.Print("Masukkan Banyaknya Kelinci : ")
    fmt.Scan(&banyak)

    for i := 0; i < banyak; i++ {
        fmt.Printf("berat Kelinci %d : ", i+1)
        fmt.Scan(&beratKelinci[i])
    }

    terkecil = beratKelinci[0]
    terbesar = beratKelinci[0]
    for i := 1; i < banyak; i++ {
        if terbesar < beratKelinci[i] {
            terbesar = beratKelinci[i]
        }

        if terkecil > beratKelinci[i] {
            terkecil = beratKelinci[i]
        }
    }

    fmt.Println("Kelinci Terbesar : ", terbesar)
    fmt.Println("Kelinci Terkecil : ", terkecil)

}

```

#### Output

![image](https://github.com/user-attachments/assets/16d3e07c-3cce-464a-af60-33ee2d60467b)

Deskripsi Program : 
Program ini digunakan untuk menentukan berat kelinci terbesar dan terkecil dari sejumlah data berat kelinci yang dimasukkan oleh pengguna. Pengguna diminta memasukkan jumlah kelinci dan berat masing-masing kelinci, yang disimpan dalam array `beratKelinci`. Nilai awal untuk berat terbesar dan terkecil diambil dari elemen pertama array. Kemudian, program menggunakan perulangan untuk membandingkan setiap nilai berat dalam array dengan nilai terbesar dan terkecil yang telah dicatat, memperbarui nilai tersebut jika ditemukan yang lebih besar atau lebih kecil. Setelah semua data selesai diproses, program mencetak berat kelinci terbesar dan terkecil ke layar. Program ini membantu pengguna melakukan analisis sederhana terhadap data berat kelinci.

### <h2>UNGUIDED 2</h2>

#### Question
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

**Masukan** terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat 𝑥 dan 𝑦. Bilangan 𝑥 menyatakan banyaknya ikan yang akan dijual, sedangkan 𝑦 adalah banyaknya ikan yang akan dimasukkan ke dalam wadah. Baris kedua terdiri dari sejumlah 𝑥 bilangan riil yang menyatakan berat ikan yang akan dijual.

**Keluaran** terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai 𝑥 dan 𝑦, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### Source Code

```go
package main

import (
    "fmt"
)

func validasi(x, y int) bool {
    return x > 0 && y > 0
}

func inputBerat(x int) []float64 {
    berat := make([]float64, x)
    for i := 0; i < x; i++ {
        fmt.Printf("Berat ikan ke-%d : ", i+1)
        fmt.Scan(&berat[i])
    }
    return berat
}

func hitungBerat(berat []float64, x, y int) {
    fmt.Println("\nTotal Berat dan Rata Rata :")
    for i := 0; i < x; i += y {
        end := i + y
        if end > x {
            end = x
        }

        var total float64
        for j := i; j < end; j++ {
            total += berat[j]
        }
        count := end - i
        fmt.Printf("Wadah %d \n", (i/y)+1)
        fmt.Printf("Total berat : %.2f \n", total)
        fmt.Printf("Rata Rata : %.2f \n", total/float64(count))
        fmt.Print("\n")
    }
}

func main() {
    var x, y int

    fmt.Print("Masukkan jumlah ikan yang dijual (x) : ")
    fmt.Scan(&x)
    fmt.Print("Masukkan jumlah ikan per wadah (y) : ")
    fmt.Scan(&y)

    if !validasi(x, y) {
        fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
        return
    }

    berat := inputBerat(x)
    hitungBerat(berat, x, y)
}

```

#### Output

![image](https://github.com/user-attachments/assets/95f0b880-76be-4ba6-9964-8aac8afa9a28)

Deskripsi Program :
Program ini digunakan untuk menghitung total dan rata-rata berat ikan yang dibagi ke dalam beberapa wadah berdasarkan kapasitas yang ditentukan. Program dimulai dengan meminta input jumlah ikan yang dijual (`x`) dan kapasitas ikan per wadah (`y`). Input divalidasi menggunakan fungsi `validasi` untuk memastikan nilainya lebih dari nol. Program kemudian meminta pengguna memasukkan berat setiap ikan, yang disimpan dalam slice `berat`. Fungsi `hitungBerat` digunakan untuk membagi ikan ke dalam wadah secara berurutan, menghitung total dan rata-rata berat ikan untuk setiap wadah, dan mencetak hasilnya. Program ini memungkinkan pengguna untuk mengelola data berat ikan secara efisien berdasarkan kapasitas wadah yang ditentukan.

### <h2>UNGUIDED 3</h2>

#### Question
Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

Buatlah program dengan spesifikasi subprogram sebagai berikut:

```go
type arrBalita [100]float64  

func hitungMinMax(arrBerat arrBalita; bMin, bMax *float64) {  
    /* I.S. Terdefinisi array dinamis arrBerat  
       F.S. Menampilkan berat minimum dan maksimum balita */  
    ...  
}  

function rerata(arrBerat arrBalita) real {  
    /* Menghitung dan mengembalikan rata-rata berat balita dalam array */  
    ...  
}  
```
Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read):

```go
Masukkan banyak data berat balita: 4  
Masukkan berat balita ke-1: 5.3  
Masukkan berat balita ke-2: 6.2  
Masukkan berat balita ke-3: 4.1  
Masukkan berat balita ke-4: 9.9  
Berat balita minimum: 4.10 kg  
Berat balita maksimum: 9.90 kg  
Rerata berat balita: 6.38 kg  
```

#### Source Code

```go
package main

import "fmt"

type arrBalita [100]float64

func minmax(arrBerat arrBalita, bMin, bMax *float64) {
    *bMin = arrBerat[0]
    *bMax = arrBerat[0]

    for i := 1; i < len(arrBerat); i++ {
        if arrBerat[i] != 0 {
            if arrBerat[i] < *bMin {
                *bMin = arrBerat[i]
            }
            if arrBerat[i] > *bMax {
                *bMax = arrBerat[i]
            }
        }
    }
}

func ratarata(arrBerat arrBalita, n int) float64 {
    var total float64
    for i := 0; i < n; i++ {
        total += arrBerat[i]
    }
    return total / float64(n)
}

func main() {
    var balita arrBalita
    var n int
    var min, max float64

    fmt.Print("Masukkan jumlah balita : ")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan berat balita ke-%d (kg) : ", i+1)
        fmt.Scan(&balita[i])
    }

    minmax(balita, &min, &max)

    rerata := ratarata(balita, n)

    fmt.Printf("\n= Hasil Analisis Data Berat Balita =\n")
    fmt.Printf("Berat Terkecil : %.2f kg\n", min)
    fmt.Printf("Berat Terbesar : %.2f kg\n", max)
    fmt.Printf("Berat Rata-rata : %.2f kg\n", rerata)
}

```

#### Output

![image](https://github.com/user-attachments/assets/c76d396f-5120-4b72-80ad-68dfdff184df)

Deskripsi Program :
Program ini digunakan untuk menganalisis data berat badan balita, mencakup pencarian berat terkecil, terbesar, dan rata-rata. Program dimulai dengan meminta input jumlah balita dan berat badan mereka, yang disimpan dalam array `balita`. Fungsi `minmax` digunakan untuk menentukan nilai berat terkecil dan terbesar dengan membandingkan setiap nilai dalam array. Fungsi `ratarata` menghitung rata-rata berat balita dengan menjumlahkan semua nilai berat dan membaginya dengan jumlah balita. Hasil analisis berupa berat terkecil, terbesar, dan rata-rata kemudian dicetak dalam format laporan yang rapi. Program ini mempermudah pengguna untuk memahami pola distribusi berat badan balita berdasarkan data yang dimasukkan.

## <strong> Kesimpulan </strong>

## Kesimpulan

Dalam bahasa Go, kita dapat dengan mudah menemukan nilai ekstrim dari sekumpulan data menggunakan struktur data seperti array atau slice. Dengan menggunakan perulangan dan kondisi sederhana, kita dapat menentukan nilai maksimum dan minimum yang diperlukan dalam berbagai aplikasi pemrograman

- Nilai ekstrim merujuk pada nilai maksimum dan minimum dalam sekumpulan data, yang penting untuk analisis dan pengambilan keputusan dalam pemrograman.
- Dalam bahasa Go, nilai ekstrim dapat dengan mudah ditemukan menggunakan perulangan untuk membandingkan elemen-elemen dalam array atau slice, menyediakan cara yang efisien untuk menganalisis data numerik.

## <strong> Referensi </strong>

#### [1] Nugroho, A., & Wibowo, S. (2022). Implementasi Algoritma Pencarian Nilai Ekstrim pada Data Numerik Menggunakan Bahasa Pemrograman Go. Jurnal Teknologi dan Sistem Komputer, 10(3), 123-130.

#### [2] Santoso, B. (2021). Belajar Dasar Pemrograman Golang. Ruang Developer. Diakses dari https://blog.ruangdeveloper.com/golang-math-operation/

#### [3] Yulianto, D. (2020). Spesifikasi Bahasa Pemrograman Go. Golang Indonesia. Diakses dari https://golang-id.org/ref/spec/
