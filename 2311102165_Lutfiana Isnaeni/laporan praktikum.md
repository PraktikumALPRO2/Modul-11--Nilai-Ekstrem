<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XI </strong></h2>
<h2 align="center"><strong> NILAI EKSTRIM </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Lutfiana Isnaeni L /2311102165<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh S.Kom., M.Kom. 
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

## Dasar Teori

**Nilai Ekstrim**

Nilai ekstrim dalam pemrograman dan algoritma merujuk pada proses untuk menemukan elemen dengan nilai terkecil (minimum) dan terbesar (maksimum) dalam sebuah kumpulan data. Konsep ini sangat berguna dalam analisis data karena membantu mengidentifikasi batas-batas signifikan pada data numerik, seperti array atau struktur data lainnya.

**Konsep Utama**

1. Nilai Minimum:
   
    Merupakan elemen terkecil dalam kumpulan data. Untuk menemukannya, setiap elemen dibandingkan secara berurutan, dan elemen dengan nilai terkecil disimpan sementara sebagai nilai minimum.

2. Nilai Maksimum:
   
    Merupakan elemen terbesar dalam kumpulan data. Proses pencariannya serupa dengan nilai minimum, namun elemen dengan nilai terbesar disimpan sementara sebagai nilai maksimum.

 **Implementasi Algoritma** 
 
   Algoritma pencarian nilai ekstrim biasanya mengikuti langkah-langkah berikut:
 
1.	Inisialisasi nilai awal untuk variabel `min` dan `max`.
   
2.	Iterasi semua elemen dalam data:
   
    •	Jika elemen lebih kecil dari `min`, maka perbarui nilai `min`.
    •	Jika elemen lebih besar dari `max`, maka perbarui nilai `max`.

3.	Setelah semua elemen selesai diproses, `min` akan menyimpan nilai terkecil, dan `max` akan menyimpan nilai terbesar.

## Unguided
### 1.  Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan rill berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
### Source Code:
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	// Cek apakah jumlah anak kelinci valid
	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	berat := make([]float64, n) // Array untuk menampung berat anak kelinci
	fmt.Println("Masukkan berat anak kelinci:")

	// Inisialisasi nilai ekstrem
	minBerat := math.MaxFloat64
	maxBerat := -math.MaxFloat64

	// Input berat dan cari nilai minimum dan maksimum
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		// Update nilai minimum dan maksimum
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	// Output hasil
	fmt.Printf("Berat kelinci terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxBerat)
}
```
#### Outpout Program
![image](https://github.com/user-attachments/assets/bc86dcbe-f18c-42b0-9931-8e7cf596200b)

### Deskripsi Program
Program ini digunakan untuk mempermudah pencatatan berat anak kelinci yang akan dijual ke pasar. Pengguna dapat menginput jumlah kelinci (maksimal 1000) beserta berat masing-masing, dan program akan secara otomatis menghitung berat terkecil serta terbesar dari data yang dimasukkan. Program ini membantu peternak atau pihak terkait dalam menganalisis data berat kelinci secara akurat, menampilkan hasil akhir berupa berat terkecil dan terbesar dalam format dua angka desimal.

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program Ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (Jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
### Source Code:
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Input jumlah wadah (x) dan jumlah ikan per wadah (y)
	fmt.Print("Masukkan jumlah wadah (x): ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Jumlah wadah dan ikan per wadah harus lebih dari 0 dan jumlah wadah maksimal 1000.")
		return
	}

	// Input berat ikan
	berat := make([]float64, x*y) // Array untuk menampung berat ikan
	fmt.Println("Masukkan berat ikan (sejumlah x * y):")
	for i := 0; i < x*y; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung total berat per wadah
	totalBeratPerWadah := make([]float64, x)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			totalBeratPerWadah[i] += berat[i*y+j]
		}
	}

	// Hitung berat rata-rata per wadah
	var totalBeratSemua float64
	for i := 0; i < x; i++ {
		totalBeratSemua += totalBeratPerWadah[i]
	}
	rataRataBerat := totalBeratSemua / float64(x)

	// Output hasil
	fmt.Println("\nHasil:")
	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < x; i++ {
		fmt.Printf("Wadah %d: %.2f\n", i+1, totalBeratPerWadah[i])
	}
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRataBerat)
}
```
#### Outpout Program
![image](https://github.com/user-attachments/assets/719a7865-0b6c-4a56-87f2-6d1a0b9c68b5)

### Deskripsi Program
Program di atas digunakan untuk menghitung total berat ikan di setiap wadah dan berat rata-rata ikan dari semua wadah berdasarkan data yang dimasukkan oleh pengguna. Program memungkinkan pengguna untuk memasukkan jumlah wadah (x), jumlah ikan per wadah (y), dan berat masing-masing ikan. Program akan secara otomatis mengelompokkan berat ikan ke dalam wadah, menghitung total berat setiap wadah, serta menentukan rata-rata berat ikan dari keseluruhan wadah. Hasil akhirnya berupa total berat ikan di setiap wadah dan rata-rata berat ikan dari semua wadah, yang ditampilkan dalam format desimal untuk memastikan ketepatan. Program ini mampu menangani hingga 1000 wadah.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:
![image](https://github.com/user-attachments/assets/63d25706-de7a-46fb-b80f-2c1ad6217dbb)
![image](https://github.com/user-attachments/assets/1fd28bcc-7171-42e9-8db3-628bb3ae3529)

### Source Code:
```go
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai minimum dan maksimum
func hitungMinMax(arrBalita []float64) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64

	for _, berat := range arrBalita {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}
	return min, max
}

// Fungsi untuk menghitung rata-rata
func hitungRataRata(arrBalita []float64) float64 {
	total := 0.0
	for _, berat := range arrBalita {
		total += berat
	}
	return total / float64(len(arrBalita))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	arrBalita := make([]float64, n)

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBalita[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	min, max := hitungMinMax(arrBalita)
	rataRata := hitungRataRata(arrBalita)

	// Tampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
```
#### Outpout Program
![image](https://github.com/user-attachments/assets/78b24bc6-aa59-4b11-bf48-a5c73cd36781)
### Deskripsi Program
Program di atas digunakan untuk membantu Posyandu dalam mencatat dan menganalisis data berat balita untuk menentukan berat terkecil, terbesar, dan rata-rata. Program ini memungkinkan pengguna memasukkan jumlah data hingga 100 balita beserta berat masing-masing. Program secara otomatis menghitung berat balita terendah dan tertinggi melalui fungsi khusus, serta menghitung rata-rata berat dengan menjumlahkan seluruh data berat lalu membaginya dengan jumlah balita.










