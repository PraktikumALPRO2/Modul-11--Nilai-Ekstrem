
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
  Mutia Rani Zahra Meilani
  <br>
  2311102182
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

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN NILAI EKSTRIM</strong></span>
<br>
Nilai ekstrem adalah nilai tertinggi (maksimum) atau terendah (minimum) dalam suatu himpunan data. Dalam analisis data, menemukan nilai ekstrem sangat penting untuk berbagai aplikasi, seperti statistik, pengolahan sinyal, dan machine learning.

- Nilai Maksimum: Nilai tertinggi dalam kumpulan data.
- Nilai Minimum: Nilai terendah dalam kumpulan data.

<span style="font-size:16px"><strong> ── ALGORITMA PENCARIAN NILAI EKSTRIM</strong></span>
<br>
Pencarian nilai ekstrem dapat dilakukan dengan menggunakan algoritma sederhana yang melibatkan iterasi melalui elemen-elemen dalam struktur data seperti array atau slice. Berikut adalah langkah-langkah umum untuk menemukan nilai maksimum dan minimum:

- **Inisialisasi** : Tentukan variabel untuk menyimpan nilai maksimum dan minimum, biasanya diinisialisasi dengan elemen pertama dari kumpulan data.
- **Iterasi** : Lakukan loop melalui setiap elemen dalam kumpulan data.
- **Perbandingan** : Bandingkan setiap elemen dengan nilai maksimum dan minimum yang saat ini disimpan, dan perbarui jika diperlukan.

Contoh Penggunaan:
```go
package main

import (
    "fmt"
)

func findExtremes(numbers []int) (int, int) {
    max := numbers[0]
    min := numbers[0]
    
    for _, num := range numbers {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    
    return max, min
}

func main() {
    data := []int{3, 5, 1, 8, 2}
    max, min := findExtremes(data)
    fmt.Printf("Nilai maksimum : %d, Nilai minimum : %d\n", max, min)
}
```

Penjelasan Kode:

- Fungsi 'findExtremes' menerima slice dari integer dan mengembalikan nilai maksimum dan minimum.
- Variabel 'max' dan 'min' diinisialisasi dengan elemen pertama dari slice.
- Loop digunakan untuk membandingkan setiap elemen dengan 'max' dan 'min', memperbarui keduanya sesuai kebutuhan.


## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case

Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat 𝑁 yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya 𝑁 bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### Source Code

```go
package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64
	var terkecil, terbesar float64

	fmt.Print("Masukkan banyaknya kelinci : ")
	fmt.Scan(&n);

	for i := 0; i < n; i++ {
    fmt.Printf("Masukan berat kelinci ke-%d : ", i+1)
    fmt.Scan(&berat[i])
  }

	terkecil = berat[0]
	terbesar = berat[0]
	for i := 1; i < n; i++ {
		if terbesar  < berat[i] {
			terbesar = berat[i]
		}

		if terkecil > berat[i] {
      terkecil = berat[i]
    }
	}

	fmt.Println("Kelinci berat terbesar : ", terbesar)
	fmt.Println("Kelinci berat terkecil : ", terkecil)

}
```

#### Output
![image](https://github.com/user-attachments/assets/ef477f08-807a-48bb-bc9c-373d56179a27)

#### Deskripsi Program :
Program ini digunakan untuk menghitung dan menentukan berat kelinci terminimum dan termaksimum dari sejumlah kelinci yang diinputkan pengguna. Program meminta pengguna memasukkan jumlah kelinci, kemudian meminta berat setiap kelinci, selanjutnya menghitung dan menampilkan berat kelinci terkecil dan terbesar dari kumpulan data yang dimasukkan.

### ── Unguided 2

#### Study Case

Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat 𝑥 dan 𝑦. Bilangan 𝑥 menyatakan banyaknya ikan yang akan dijual, sedangkan 𝑦 adalah banyaknya ikan yang akan dimasukkan ke dalam wadah. Baris kedua terdiri dari sejumlah 𝑥 bilangan riil yang menyatakan berat ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai 𝑥 dan 𝑦, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### Source Code

```go
package main

import (
	"fmt"
)

func validasi(x, y int) bool {
	return x > 0 && y > 0
}

func input_berat(x int) []float64 {
	berat := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d : ", i+1)
		fmt.Scan(&berat[i])
	}
	return berat
}

func hitung_berat(berat []float64, x, y int) {
	fmt.Println("\n= Hasil Perhitungan Total Berat dan Rata Rata =")
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
		fmt.Printf("WADAH %d \n", (i/y)+1)
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

	berat := input_berat(x)
	hitung_berat(berat, x, y)
}
```

#### Output
![image](https://github.com/user-attachments/assets/30448d23-3bec-4f12-9727-913642ec9847)

#### Deskripsi Program :
Program ini dirancang untuk menghitung total berat dan rata-rata berat ikan yang dijual dalam beberapa wadah. Program meminta pengguna memasukkan jumlah total ikan dan jumlah ikan per wadah, kemudian meminta input berat setiap ikan. Selanjutnya program menghitung dan menampilkan total berat serta rata-rata berat ikan untuk setiap wadah yang telah ditentukan.

### ── Unguided 3

#### Study Case

Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

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
        fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
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
![image](https://github.com/user-attachments/assets/47caebfa-9cff-4600-88ac-f465fde6c630)

#### Deskripsi Program :
Program ini bertujuan untuk menganalisis data berat balita dengan menghitung berat terkecil, berat terbesar, dan berat rata-rata. Program meminta pengguna memasukkan jumlah balita dan berat masing-masing balita, kemudian menggunakan fungsi khusus untuk menentukan berat minimum dan maksimum serta menghitung rata-rata berat balita yang diinputkan.


## <strong> REFERENSI </strong>

#### [1] Santoso, A., & Prabowo, H. (2020). "Analisis Algoritma Pencarian Nilai Ekstrem pada Data Besar". Jurnal Teknologi Informasi dan Komunikasi.
#### [2] Knuth, D. E. (1997). "The Art of Computer Programming". Volume 1: Fundamental Algorithms.
#### [3] Golang.org Documentation. Retrieved from https://golang.org/doc/.
