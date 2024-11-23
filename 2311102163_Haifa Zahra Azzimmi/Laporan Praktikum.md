![image](https://github.com/user-attachments/assets/40573bda-1f80-4a7e-b740-23980a173ac6)<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL 11</strong></h2>
<h2 align="center"><strong> NILAI EKSTREM </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Haifa Zahra Azzzimmi / 2311102163<br>
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

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan rill berikutnya adalah berat dari anak kelinci yang akan dijual.

Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	beratKelinci := make([]float64, N)
	fmt.Printf("Masukkan berat %d anak kelinci:\n", N)
	for i := 0; i < N; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	for _, berat := range beratKelinci {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}

```
### Output:
![image](https://github.com/user-attachments/assets/4c50c7b0-b043-4140-85e1-2d4c9bfed387)

### Full code Screenshot:
![code  7](https://github.com/user-attachments/assets/039d0f6a-2e5d-433d-8d83-4f39a870d1a2)

### Deskripsi Program : 
Program ini dirancang untuk mencari berat terkecil dan terbesar di antara anak-anak kelinci. Pertama-tama, program meminta pengguna untuk memasukkan jumlah anak kelinci. Jika jumlahnya valid (antara 1 hingga 1000), program kemudian meminta berat masing-masing kelinci tersebut. Setelah itu, program menggunakan loop untuk memeriksa setiap berat guna menemukan yang terkecil dan terbesar. Pada akhirnya, program akan mencetak berat terkecil dan terbesar dengan dua angka di belakang koma.

### 2.  Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program Ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (Jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus valid (x <= 1000, y > 0).")
		return
	}

	fmt.Printf("Masukkan berat %d ikan:\n", x)
	beratIkan := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalBerat float64
	wadah := []float64{}
	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]
		if (i+1)%y == 0 || i == x-1 {
			wadah = append(wadah, totalBerat)
			totalBerat = 0
		}
	}

	var totalWadah float64
	for _, berat := range wadah {
		totalWadah += berat
	}
	rataRata := totalWadah / float64(len(wadah))

	fmt.Println("Total berat di setiap wadah:")
	for _, berat := range wadah {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}

```
### Output:
![image](https://github.com/user-attachments/assets/c3205675-42b8-49d4-9769-e04357fdbc7e)

### Full code Screenshot:
![code  8](https://github.com/user-attachments/assets/6676852c-1e6b-4a2e-9fb5-41e631cac56d)

### Deskripsi Program : 
Program ini bertujuan untuk menghitung total berat ikan dalam beberapa wadah dan kemudian mencari rata-rata berat per wadah. Pertama, program meminta pengguna untuk memasukkan jumlah ikan dan kapasitas wadah. Setelah itu, pengguna diminta untuk memasukkan berat masing-masing ikan. Program kemudian mengelompokkan ikan ke dalam wadah sesuai kapasitasnya, menghitung total berat untuk setiap wadah, dan akhirnya, menghitung serta mencetak rata-rata berat per wadah.
 
### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

Buatlah program dengan spesifikasi subprogram sebagai berikut:
![image](https://github.com/user-attachments/assets/0846cc49-9d9d-4357-b7e3-7c4707cd9fef)
![image](https://github.com/user-attachments/assets/330f9db4-44b9-4a9f-8c01-2ebb6446e477)

### Source Code :
```go
//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
)

type arrBalita [100]float64 // Definisi array untuk data berat balita

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 hingga 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)
	rerataBerat := rerata(berat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBerat)
}

```
### Output:
![image](https://github.com/user-attachments/assets/4ea28b0f-0d8d-4bb6-ba56-f3baa563766c)

### Full code Screenshot:
![code  9](https://github.com/user-attachments/assets/e4cf4847-9889-49c3-8d93-208b540b7f54)

### Deskripsi Program : 
Program ini menghitung berat minimum, maksimum, dan rata-rata dari data berat balita. Pertama, program meminta pengguna untuk memasukkan jumlah data berat balita. Jika jumlahnya valid (antara 1 hingga 100), program kemudian meminta berat masing-masing balita. Dengan fungsi hitungMinMax, program mencari berat terkecil dan terbesar. Fungsi rerata digunakan untuk menghitung rata-rata berat balita. Pada akhirnya, program mencetak hasil berat minimum, maksimum, dan rata-rata.

