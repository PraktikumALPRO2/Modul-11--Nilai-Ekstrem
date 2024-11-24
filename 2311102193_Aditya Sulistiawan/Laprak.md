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
 Aditya Sulistiawan
  <br>
  2311102193
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
Dalam pemrograman, nilai ekstrem merujuk pada nilai terkecil (minimum) dan terbesar (maksimum) dalam sebuah kumpulan data. Dalam konteks bahasa Go (Golang), nilai ekstrem sering dicari menggunakan struktur data seperti array atau slice, dan prosesnya melibatkan iterasi untuk membandingkan elemen satu per satu.

### <strong> Cara Mencari Nilai Ekstrem di Go:
1. Inisialisasi Nilai Awal:
	Untuk mencari nilai minimum, inisialisasi dengan nilai maksimum (math.MaxFloat64 untuk float64 atau math.MaxInt untuk integer).
	Untuk mencari nilai maksimum, inisialisasi dengan nilai minimum (-math.MaxFloat64 atau math.MinInt).

2. Iterasi Data:
	Bandingkan setiap elemen dalam array/slice dengan nilai saat ini.
	Perbarui nilai minimum/maksimum jika elemen saat ini lebih kecil/besar dari nilai yang sudah tersimpan.

3. Efisiensi: Pendekatan ini memiliki kompleksitas waktu O(n), karena setiap elemen diperiksa satu kali.
  
### <strong> Implementasi:
```go
package main

import (
	"fmt"
	"math"
)

func cariMinMax(data []float64) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64
	for _, val := range data {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}

func main() {
	data := []float64{2.5, 3.1, 7.4, 1.2, 4.6}
	min, max := cariMinMax(data)
	fmt.Printf("Nilai minimum: %.2f, Nilai maksimum: %.2f\n", min, max)
}
```

# <strong> Unguided </strong>
## Unguided - 1
### Study Case
**1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.**

_Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual._                                               


_Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar._ 


### Source Code
```go
//Aditya Sulistiawan
//2311102193

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk mendapatkan berat kelinci dari pengguna
func ambilBeratKelinci(jumlah int) []float64 {
	berat := make([]float64, jumlah)
	fmt.Println("Masukkan berat kelinci:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	return berat
}

// Fungsi untuk mencari berat terkecil dan terbesar
func cariMinMax(berat []float64) (float64, float64) {
	beratTerkecil := math.MaxFloat64
	beratTerbesar := -math.MaxFloat64

	for _, b := range berat {
		if b < beratTerkecil {
			beratTerkecil = b
		}
		if b > beratTerbesar {
			beratTerbesar = b
		}
	}

	return beratTerkecil, beratTerbesar
}

func main() {
	// Input jumlah anak kelinci
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&jumlahKelinci)

	// Validasi jumlah anak kelinci
	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	// Ambil berat kelinci dan hitung berat terkecil dan terbesar
	berat := ambilBeratKelinci(jumlahKelinci)
	beratTerkecil, beratTerbesar := cariMinMax(berat)

	// Tampilkan hasil
	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}
```

### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/ed9dbd6c-73bc-487f-bf10-097c887e0291)


### Screenshot Output
![Gui](https://github.com/user-attachments/assets/0c04cc62-ca3b-4c27-8995-744447fa3f28)


### Deskripsi Program
Program ini digunakan untuk menentukan berat terkecil dan terbesar dari sejumlah anak kelinci. Pengguna memasukkan jumlah anak kelinci dan berat masing-masing kelinci, kemudian program menghitung berat terkecil dan terbesar dari data yang diberikan.

### Algoritma Program
1. Input Jumlah Kelinci: Minta pengguna memasukkan jumlah anak kelinci (N).
2. Validasi Input: Periksa apakah N berada dalam rentang 1 hingga 1000.
3. Input Berat Kelinci: Minta pengguna memasukkan berat tiap kelinci.
4. Cari Berat Terkecil dan Terbesar
5. Output Hasil: Tampilkan berat terkecil dan terbesar ke layar.

Cara kerja program nya adalah dimulai dengan meminta pengguna untuk memasukkan jumlah kelinci yang akan diproses. Apabila jumlah yang dimasukkan tidak valid (kurang dari 1 atau lebih dari 1000), program akan menunjukkan pesan kesalahan dan menghentikan prosesnya. Setelah jumlah yang valid telah diberikan, pengguna akan diminta untuk memasukkan berat setiap kelinci satu per satu. Program kemudian mengolah data tersebut untuk menentukan berat paling ringan dan paling berat dengan membandingkan semua berat yang diinputkan. Terakhir, program menampilkan hasil berupa berat paling ringan dan paling berat di layar sebagai output.

  ## Unguided - 2
### Study Case
**2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.**

_Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah, Baris kedua terdiri dari sejumlah x bilangan rill yang menyatakan banyaknya ikan yang akan dijual._                                                                                  


_Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan niil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesual urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah._


### Source Code
```go
//Aditya Sulistiawan
//2311102193

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
### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/1ff84dfc-0778-4949-abad-5365f68d6578)



### Screenshot Output
![1](https://github.com/user-attachments/assets/487bb98f-9e26-4aa8-8ef0-a82b0bcc8bcc)


### Deskripsi Program
Program ini menghitung jumlah total berat ikan di setiap wadah dan rata-rata berat per wadah. Pengguna memasukkan jumlah ikan, kapasitas wadah, dan berat dari masing-masing ikan. Program selanjutnya akan membagi ikan ke dalam wadah sesuai dengan kapasitas yang telah ditentukan, menghitung total berat di setiap wadah, serta menampilkan rata-rata berat dari seluruh wadah.

**Algoritma Program**
1. Input Jumlah Ikan dan Kapasitas Wadah**: Minta pengguna memasukkan jumlah ikan (`x`) dan kapasitas setiap wadah (`y`).
2. Validasi Input**: Periksa apakah jumlah ikan valid (1 ≤ x ≤ 1000) dan kapasitas wadah lebih dari 0.
3. Input Berat Ikan**: Simpan berat masing-masing ikan ke dalam array.
4. Hitung Total Berat Per Wadah**:
   - Tambahkan berat ikan secara berurutan.
   - Saat jumlah ikan mencapai kapasitas wadah (`y`), simpan total berat ke daftar total wadah dan ulangi proses untuk ikan berikutnya.
5. Hitung Rata-rata Berat Wadah**: Jumlahkan semua berat dalam daftar total wadah dan bagi dengan jumlah wadah.
6. Tampilkan Hasil**: Cetak total berat ikan per wadah dan rata-rata berat wadah.

cara kerja program nya yaitu program ini nantinya meminta jumlah ikan dan kapasitas wadah, lalu memvalidasi inputnya. Pengguna memasukkan berat tiap ikan, yang kemudian dikelompokkan ke dalam wadah sesuai kapasitas. Program menghitung total berat tiap wadah dan rata-rata berat semua wadah, lalu menampilkan hasilnya gitu deh kak.	

  ## Unguided - 3
### Study Case
**3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.**


### Source Code
```go
///Aditya Sulistiawan
///2311102193

package main

import (
	"fmt"
)

// Definisi tipe array untuk menyimpan data berat anak
type dataBeratAnak [100]float64

// Fungsi untuk menentukan berat minimum dan maksimum
func cariMinMax(beratAnak dataBeratAnak, jumlah int, beratMin, beratMax *float64) {
	*beratMin = beratAnak[0]
	*beratMax = beratAnak[0]
	for idx := 1; idx < jumlah; idx++ {
		if beratAnak[idx] < *beratMin {
			*beratMin = beratAnak[idx]
		}
		if beratAnak[idx] > *beratMax {
			*beratMax = beratAnak[idx]
		}
	}
}

// Fungsi untuk menghitung nilai rata-rata berat
func hitungRataRata(beratAnak dataBeratAnak, jumlah int) float64 {
	totalBerat := 0.0
	for idx := 0; idx < jumlah; idx++ {
		totalBerat += beratAnak[idx]
	}
	return totalBerat / float64(jumlah)
}

func main() {
	var jumlahAnak int
	var berat dataBeratAnak
	var beratMin, beratMax float64

	// Input jumlah anak
	fmt.Print("Masukkan jumlah data berat anak: ")
	fmt.Scan(&jumlahAnak)

	// Validasi jumlah data
	if jumlahAnak <= 0 || jumlahAnak > 100 {
		fmt.Println("Jumlah data berat anak harus antara 1 dan 100.")
		return
	}

	// Input data berat masing-masing anak
	for i := 0; i < jumlahAnak; i++ {
		fmt.Printf("Masukkan berat anak ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung berat minimum, maksimum, dan rata-rata
	cariMinMax(berat, jumlahAnak, &beratMin, &beratMax)
	rataRata := hitungRataRata(berat, jumlahAnak)

	// Output hasil perhitungan
	fmt.Printf("Berat anak minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat anak maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat anak: %.2f kg\n", rataRata)
}
```
### Screenshoot Source Code
![carbon](https://github.com/user-attachments/assets/e705cf04-3e2c-4bd3-842b-1af491e03b6e)


### Screenshot Output
![gui](https://github.com/user-attachments/assets/1c6dc5ea-efa3-4791-929e-a984bc757d60)


### Deskripsi Program
Program ini digunakan untuk menghitung berat minimum, maksimum, dan rata-rata dari sejumlah data berat anak yang dimasukkan oleh pengguna.

**Algoritma Program**
1. Input jumlah anak dan validasi bahwa nilainya berada dalam rentang 1–100.
2. Input berat masing-masing anak.
3. Cari berat terkecil dan terbesar dengan membandingkan setiap berat yang dimasukkan.
4. Hitung rata-rata berat dengan menjumlahkan semua berat lalu membagi dengan jumlah anak.
5. Tampilkan berat minimum, maksimum, dan rata-rata.

cara kerja program : Program dimulai dengan meminta jumlah data berat anak dan memvalidasi apakah jumlahnya berada dalam rentang yang diperbolehkan. Setelah itu, pengguna memasukkan berat masing-masing anak satu per satu. Program menghitung berat minimum dan maksimum dengan membandingkan semua data berat, kemudian menghitung rata-rata berat anak. Hasil berupa berat minimum, maksimum, dan rata-rata ditampilkan ke layar.

  ## Referensi
[1] The GO Programming Language specification - The GO Programming language. (n.d.).

[2] Documentation - The Go Programming Language. (n.d.).
