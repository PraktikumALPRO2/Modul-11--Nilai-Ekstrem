<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL XI</strong></h2>
<h2 align="center"><strong> NILAI EKSTREM</strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Ria Wulandari / 2311102173<br>
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

## Dasar Teori
Nilai ekstrem mengacu pada elemen terbesar (maksimum) atau terkecil (minimum) dalam sekumpulan data. Pencarian nilai ekstrem didasarkan pada konsep iterasi dan pembandingan. Berikut adalah beberapa poin penting:
### Struktur Data yang Digunakan
Biasanya, nilai ekstrem dicari dalam struktur data seperti array atau slice. Pemilihan struktur ini memengaruhi cara iterasi dan efisiensi pencarian.
### Algoritma Dasar
- Mulai dengan menetapkan elemen pertama sebagai kandidat nilai ekstrem.
- Iterasi melalui elemen lain satu per satu.
- Bandingkan setiap elemen dengan kandidat nilai ekstrem, lalu perbarui jika ditemukan nilai yang lebih besar (maksimum) atau lebih kecil (minimum).
### Kompleksitas Algoritma
- Kompleksitas waktu: O(n), karena setiap elemen diperiksa satu kali.
- Kompleksitas ruang: O(1), karena hanya membutuhkan variabel tambahan untuk menyimpan nilai ekstrem.
### Batasan Tipe Data
Nilai ekstrem dipengaruhi oleh batasan tipe data seperti `int`, `float`, atau tipe lainnya di Go. Hal ini perlu diperhatikan untuk menghindari overflow atau underflow.
### Penerapan Praktis
Pencarian nilai ekstrem digunakan dalam banyak kasus, seperti pengolahan data statistik, optimisasi algoritma, atau analisis kondisi ekstrem dalam dataset.
## Unguided
### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### Source code :
``` go
package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&jumlahKelinci)

	// Validasi jumlah anak kelinci
	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	beratKelinci := make([]float64, jumlahKelinci) // Slice untuk berat kelinci
	fmt.Println("Masukkan berat setiap anak kelinci:")

	// Nilai ekstrem untuk berat
	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	// Loop untuk input berat dan cari nilai ekstrem
	for indeks := 0; indeks < jumlahKelinci; indeks++ {
		fmt.Printf("Berat kelinci ke-%d: ", indeks+1)
		fmt.Scan(&beratKelinci[indeks])

		// Pembaruan nilai ekstrem
		if beratKelinci[indeks] < beratTerkecil {
			beratTerkecil = beratKelinci[indeks]
		}
		if beratKelinci[indeks] > beratTerbesar {
			beratTerbesar = beratKelinci[indeks]
		}
	}

	// Menampilkan hasil
	fmt.Printf("Berat kelinci terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", beratTerbesar)
}
```
#### Output
![image](https://github.com/user-attachments/assets/f9971af9-285a-4715-a5bd-d53a6d3a6fb6)
#### Deskripsi
Program ini meminta pengguna untuk memasukkan jumlah anak kelinci (dengan batas valid antara 1 hingga 1000). Setelah itu, pengguna diminta memasukkan berat masing-masing anak kelinci satu per satu. Program kemudian akan menentukan berat terkecil dan terbesar dari data yang dimasukkan, lalu menampilkan hasilnya kepada pengguna.
### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program Ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (Jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### Source code :
```go
package main

import (
	"fmt"
)

func main() {
	var jumlahWadah, ikanPerWadah int

	// Input jumlah wadah dan jumlah ikan per wadah
	fmt.Print("Masukkan jumlah wadah: ")
	fmt.Scan(&jumlahWadah)

	fmt.Print("Masukkan jumlah ikan per wadah: ")
	fmt.Scan(&ikanPerWadah)

	// Validasi input
	if jumlahWadah <= 0 || jumlahWadah > 1000 || ikanPerWadah <= 0 {
		fmt.Println("Jumlah wadah harus antara 1 hingga 1000, dan jumlah ikan per wadah harus lebih dari 0.")
		return
	}

	// Input berat ikan
	beratIkan := make([]float64, jumlahWadah*ikanPerWadah) // Slice untuk menyimpan berat ikan
	fmt.Println("Masukkan berat setiap ikan (jumlah total: jumlah wadah x jumlah ikan per wadah):")
	for indeks := 0; indeks < jumlahWadah*ikanPerWadah; indeks++ {
		fmt.Printf("Berat ikan ke-%d: ", indeks+1)
		fmt.Scan(&beratIkan[indeks])
	}

	// Menghitung total berat per wadah
	totalBerat := make([]float64, jumlahWadah)
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		for ikan := 0; ikan < ikanPerWadah; ikan++ {
			totalBerat[wadah] += beratIkan[wadah*ikanPerWadah+ikan]
		}
	}

	// Menghitung rata-rata berat per wadah
	var totalKeseluruhanBerat float64
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		totalKeseluruhanBerat += totalBerat[wadah]
	}
	rataRata := totalKeseluruhanBerat / float64(jumlahWadah)

	// Menampilkan hasil
	fmt.Println("\nHasil:")
	fmt.Println("Total berat ikan di setiap wadah:")
	for wadah := 0; wadah < jumlahWadah; wadah++ {
		fmt.Printf("Wadah %d: %.2f\n", wadah+1, totalBerat[wadah])
	}
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRata)
}
```
#### Output
![image](https://github.com/user-attachments/assets/932df8b2-d4ad-42d6-a2a8-64e712dabcbf)
#### Deskripsi
Program ini meminta pengguna untuk memasukkan jumlah wadah dan jumlah ikan per wadah, dengan validasi bahwa jumlah wadah harus antara 1 hingga 1000 dan jumlah ikan lebih dari 0. Selanjutnya, pengguna diminta memasukkan berat setiap ikan. Program kemudian menghitung total berat ikan di setiap wadah dengan menjumlahkan berat ikan di wadah tersebut. Setelah itu, program menghitung berat rata-rata ikan per wadah dengan membagi total keseluruhan berat ikan dengan jumlah wadah. Terakhir, program menampilkan total berat ikan di setiap wadah dan berat rata-rata ikan per wadah.
### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:
![image](https://github.com/user-attachments/assets/0b687e63-c2fb-40f0-824b-6b62d4403530)
![image](https://github.com/user-attachments/assets/dd258770-85d4-4157-b284-63c1c7afa7f0)
#### Source code :
```go
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk mencari nilai ekstrem (minimum dan maksimum)
func cariEkstrem(beratBalita []float64) (float64, float64) {
	beratMinimal := math.MaxFloat64
	beratMaksimal := -math.MaxFloat64

	for _, berat := range beratBalita {
		if berat < beratMinimal {
			beratMinimal = berat
		}
		if berat > beratMaksimal {
			beratMaksimal = berat
		}
	}
	return beratMinimal, beratMaksimal
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRata(beratBalita []float64) float64 {
	totalBerat := 0.0
	for _, berat := range beratBalita {
		totalBerat += berat
	}
	return totalBerat / float64(len(beratBalita))
}

func main() {
	var jumlahData int
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahData)

	// Validasi jumlah data
	if jumlahData <= 0 || jumlahData > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	beratBalita := make([]float64, jumlahData)

	// Input berat balita
	for indeks := 0; indeks < jumlahData; indeks++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", indeks+1)
		fmt.Scan(&beratBalita[indeks])
	}

	// Menghitung nilai ekstrem dan rata-rata
	beratMin, beratMax := cariEkstrem(beratBalita)
	rataBerat := hitungRata(beratBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita terendah: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita tertinggi: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataBerat)
}
```
#### Output
![image](https://github.com/user-attachments/assets/d5977fed-6ed4-4c05-bbd2-133eef5baa21)
#### Deskripsi
Program ini meminta pengguna untuk memasukkan jumlah data berat balita (antara 1 hingga 100). Setelah itu, pengguna diminta untuk memasukkan berat setiap balita. Program kemudian menghitung nilai ekstrem (berat terendah dan tertinggi) dengan mencari berat yang paling kecil dan paling besar dari data yang dimasukkan. Selain itu, program juga menghitung rata-rata berat balita dengan menjumlahkan semua berat dan membaginya dengan jumlah data. Hasilnya, program menampilkan berat balita terendah, tertinggi, dan rata-rata berat balita.
