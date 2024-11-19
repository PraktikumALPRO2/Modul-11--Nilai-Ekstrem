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
  <strong>Nilai Ekstrem</strong>
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
 Rangga Pradarrell Fathi
  <br>
  2311102200
  <br>
  IF-11-05
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

<strong><h2>Definisi</h2></strong>
Pencarian nilai ekstrim adalah algoritma untuk menelusuri dan membandingkan setiap elemen dalam sebuah himpunan data guna menentukan elemen dengan nilai tertinggi (maksimum) atau nilai terendah (minimum).

### <strong> Prinsip algoritma pencarian nilai ekstrim adalah:
- Memulai dengan menetapkan elemen pertama sebagai nilai awal (sementara) untuk nilai maksimum atau minimum.
- Membandingkan setiap elemen berikutnya dengan nilai sementara.
- Memperbarui nilai maksimum atau minimum jika ditemukan elemen dengan nilai lebih tinggi (untuk maksimum) atau lebih rendah (untuk minimum).
  
### <strong> Implementasi:
Dalam pemrograman, pencarian nilai ekstrim sering kali menjadi dasar untuk membangun algoritma yang lebih kompleks, seperti sorting, optimasi, atau analisis data. Algoritma ini diimplementasikan menggunakan perulangan (loop) untuk memindai elemen-elemen dalam struktur data.

**Penggunaan:**
- Menemukan nilai tertinggi dalam data skor ujian.
- Menentukan suhu maksimum/minimum dari data cuaca.
- Mengidentifikasi harga saham tertinggi atau terendah dalam periode tertentu.

# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.**

_Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual._                                               


_Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar._ 


### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05

package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	weights := make([]float64, n)

	fmt.Println("Masukkan berat kelinci:")
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight := weights[0]
	maxWeight := weights[0]

	for _, weight := range weights {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}

```

### Screenshot Output
![image](https://github.com/user-attachments/assets/0934bcfa-5728-4594-ab02-ff9c9dc17977)

### Deskripsi Program
Program ini dirancang untuk mencari berat terkecil dan terbesar dari sejumlah anak kelinci berdasarkan input dari pengguna. Pengguna diminta untuk memasukkan jumlah anak kelinci serta berat masing-masing anak kelinci.

Algoritma Program
-Masukkan Jumlah Anak Kelinci (n):
- Inisialisasi Slice weights:
- Masukkan Berat Masing-Masing Anak Kelinci:
- Inisialisasi Berat Terkecil dan Terbesar:
- Pencarian Berat Terkecil dan Terbesar:
- Iterasi melalui seluruh elemen slice weights:
- Output Hasil

  ## Unguided - 2
### Study Case
**2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.**

_Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah, Baris kedua terdiri dari sejumlah x bilangan rill yang menyatakan banyaknya ikan yang akan dijual._                                                                                  


_Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan niil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesual urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah._


### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11-05
package main

import (
	"fmt"
)

func main() {
	// Input x dan y
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi x dan y
	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, dan kapasitas wadah harus positif.")
		return
	}

	// Input berat ikan
	weights := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	// Menghitung total berat per wadah
	totalWeights := []float64{}
	currentWeight := 0.0
	for i := 0; i < x; i++ {
		currentWeight += weights[i]
		if (i+1)%y == 0 || i == x-1 {
			totalWeights = append(totalWeights, currentWeight)
			currentWeight = 0.0
		}
	}

	// Menghitung rata-rata berat per wadah
	totalSum := 0.0
	for _, weight := range totalWeights {
		totalSum += weight
	}
	averageWeight := totalSum / float64(len(totalWeights))

	// Output hasil
	fmt.Println("Total berat ikan per wadah:")
	for _, weight := range totalWeights {
		fmt.Printf("%.2f ", weight)
	}
	fmt.Println()
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", averageWeight)
}

```

### Screenshot Output
![image](https://github.com/user-attachments/assets/1ed7a619-61c4-4ecc-af06-50e291a00da8)

### Deskripsi Program
Program ini dirancang untuk menghitung total berat ikan dalam wadah dan rata-rata berat ikan per wadah berdasarkan jumlah ikan, kapasitas wadah, dan berat masing-masing ikan yang diinputkan pengguna. 

**Algoritma Program**
- Input Jumlah Ikan (x) dan Kapasitas Wadah (y).
- Input Berat Ikan.
- Pengelompokan Berat Ikan ke Dalam Wadah.
- Menghitung Rata-rata Berat per Wadah.
- Output Hasil.

  ## Unguided - 2
### Study Case
**3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.**


### Source Code
```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import (
	"fmt"
)

// Definisi tipe array untuk berat balita
type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, min, max *float64) {
	*min = arrBerat[0]
	*max = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *min {
			*min = arrBerat[i]
		}
		if arrBerat[i] > *max {
			*max = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rata-rata berat balita
func rataRata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var min, max float64

	// Input jumlah data balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data balita harus antara 1 dan 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(berat, n, &min, &max)
	rata := rataRata(berat, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}

```

### Screenshot Output
![image](https://github.com/user-attachments/assets/8a6da666-03d3-43d9-b869-8a279772dcb7)

### Deskripsi Program
Program ini digunakan untuk menghitung berat minimum, maksimum, dan rata-rata dari sejumlah data berat balita yang dimasukkan oleh pengguna. Program ini memanfaatkan array statis (arrBalita) untuk menyimpan data berat balita dan beberapa subprogram untuk menghitung nilai-nilai yang diperlukan.

**Algoritma Program**
- Meminta pengguna memasukkan jumlah data balita (n).
- Memvalidasi bahwa n berada di antara 1 dan 100. Jika tidak, program akan berhenti dengan pesan kesalahan.
- Input Berat Balita:

  ## Referensi
[1] Knuth, D. E. (1997). The Art of Computer Programming, Volume 3: Sorting and Searching. Addison-Wesley.                                                       

[2] Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2009). Introduction to Algorithms. MIT Press.

