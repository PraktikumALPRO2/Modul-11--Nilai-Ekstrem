<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XI</strong></h2>
<h2 align="center"><strong> NILAI EKSTREM </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Andreas Besar Wibowo / 2311102198<br>
  S1 IF11-05
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

## I. Dasar Teori
### Definisi Pencarian 
Penyortiran dan pencarian merupakan operasi dasar dalam ilmu komputer dan memainkan peran penting dalam berbagai aplikasi. Algoritme penyortiran menyusun elemen dalam urutan tertentu, sementara algoritme pencarian menemukan keberadaan dan lokasi elemen target dalam kumpulan data[1].

### Definisi Pencarian Linear
Pencarian linear, juga dikenal sebagai pencarian sekuensial, adalah algoritma pencarian sederhana dan mudah digunakan untuk menemukan elemen target tertentu dalam daftar atau larik. Dalam metode pencarian ini, algoritma memeriksa setiap elemen daftar secara berurutan, mulai dari awal, hingga elemen target ditemukan atau seluruh daftar telah dilintasi[1].

### Definisi Pencarian Binary
Pencarian Biner didefinisikan sebagai algoritma pencarian yang digunakan dalam array yang diurutkan dengan membagi interval pencarian menjadi dua secara berulang. Ide pencarian biner adalah menggunakan informasi bahwa array diurutkan dan mengurangi kompleksitas waktu[2].

## II. UNGUIDED
**1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.**

_Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual._                                               


_Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar._ 

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Tipe data array untuk menyimpan berat kelinci
type Kelinci struct {
	Berat float64
}

func main() {
	// Untuk Inputan jumlah anak kelinci
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	// Mevalidasi jumlah anak kelinci
	if N < 1 || N > 1000 {
		fmt.Println("Jumlah anak kelinci tidak lebih dari 1000.")
		return
	}

	// untuk Array Kelinci
	daftarKelinci := make([]Kelinci, N)

	// Untuk Inputan berat anak-anak kelinci
	fmt.Println("Masukkan berat anak-anak kelinci:")
	for i := 0; i < N; i++ {
		var berat float64
		fmt.Printf("Berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat)
		// Memasukkan berat kelinci ke Array Kelinci
		daftarKelinci[i] = Kelinci{Berat: berat}
	}

	minKelinci := daftarKelinci[0] // untuk nilai min dengan berat kelinci pertama
	maxKelinci := daftarKelinci[0] // untuk nilai max dengan berat kelinci pertama

	// untuk mencari berat min dan max kelinci
	for _, kelinci := range daftarKelinci {
		if kelinci.Berat < minKelinci.Berat {
			minKelinci = kelinci
		}
		if kelinci.Berat > maxKelinci.Berat {
			maxKelinci = kelinci
		}
	}

	// Untuk Output Akhir
	fmt.Printf("Berat terkecil: %.2f kg\n", minKelinci.Berat)
	fmt.Printf("Berat terbesar: %.2f kg\n", maxKelinci.Berat)
}
```
#### Screenshoot Source Code
![Unguided1 carbon](https://github.com/user-attachments/assets/e00a1069-1723-45c3-a7b6-b075452b5176)

#### Screenshoot Output
![Unguided1 go](https://github.com/user-attachments/assets/d2cb8e62-d8d8-4b90-be91-832cbd86ed9f)

#### Deskripsi Program
Program ini merupakan program untuk menghitung dan menampilkan berat kelinci dari terkecil dan terbesar yang beratnya diinputkan oleh pengguna. Proses pencarian nilai ekstremnya menggunakan algoritma pencarian linear. Program ini menggunakan array bertipe data terstruktur karena menyimpan data 'Kelinci' dengan 'Berat' bertipe float64

Algortima Program :
1. Input jumlah anak kelinci dan berat pada masing-masing kelinci
2. Inisialisasi berat terkecil dan terbesar 
3. Lakukan algoritma pencarian linear untuk memeriksa dan mencari nilai terkecil dan terbesar
4. Tampilkan berat anak kelinci dari terkecil dan terbesar.

Cara kerja dari program ini yaitu user memasukkan jumlah anak kelinci dan beratnya masing-masing. Program menginisialisasi berat kelinci dan melakukan pencarian linear untuk memeriksa berat kelinci. Dalam proses ini melakukan perbandingan berat kelinci dan mencari nilai min dan max. Pada output akhir akan menampilkan berat kelinci yang terkecil dan terbesar.

**2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.**

_Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah, Baris kedua terdiri dari sejumlah x bilangan rill yang menyatakan banyaknya ikan yang akan dijual._                                                                                  


_Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan niil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesual urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah._

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

func main() {
	// Untuk jumlah ikan (x) dan jumlah ikan per wadah (y)
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah : ")
	fmt.Scan(&x, &y)

	// Mevalidasi jumlah ikan per wadah
	if x < 1 || x > 1000 || y < 1 {
		fmt.Println("Jumlah ikan harus tidak lebih dari 1000 dan jumlah ikan per wadah harus lebih dari 0")
		return
	}

	// array untuk menyimpan berat ikan
	ikan := make([]float64, x) 
	fmt.Println("Masukkan berat ikan :")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i]) 
	}

	
	// Menghitung jumlah wadah yang dibutuhkan
	totWadah := (x + y - 1) / y 
	beratPerWadah := make([]float64, totWadah) 

	// untuk mengisi array beratPerWadah dengan total berat ikan pada setiap wadah
	for i := 0; i < x; i++ {
		idxWadah := i / y 
		beratPerWadah[idxWadah] += ikan[i] 
	}

	// Output total berat ikan pada setiap wadah
	fmt.Println("Total berat ikan pada setiap wadah :")
	for i, berat := range beratPerWadah {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, berat) 
	}

	// untuk menghitung berat rata-rata ikan pada setiap wadah
	var totalBerat float64 
	for _, berat := range beratPerWadah {
		totalBerat += berat 
	}
	rataRataBerat := totalBerat / float64(totWadah) // untuk menghitung rata-rata berat

	// Output rata-rata berat ikan 
	fmt.Printf("Berat rata-rata ikan pada setiap wadah: %.2f kg\n", rataRataBerat)
}
```
#### Screenshoot Source Code
![Unguided2 carbon](https://github.com/user-attachments/assets/57712e0a-0cff-4d5d-8049-0c00bb785922)

#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/2aa2649f-b94f-4919-aa6e-f63ca851b33c)

#### Deskripsi Program
Program ini merupakan program untuk menghitung dan menampilkan berat ikan per wadah dan rata-rata berat ikan per wadah dari jumlah ikan dan berat ikan yang diinputkan oleh user. Program ini menggunakan algoritma pencarian linear. Program ini juga menggunakan array bertipe data dasar yntuk menyimpan berat ikan dengan tipe data float64

Algortima program :
1. Input jumlah ikan dan jumlah ikan per wadah
2. Validasi jumlah ikan dan kapasitas wadah tidak lebih dari 1000
3. Input berat ikan masing masing
4. Menghitung jumlah wadah yang dibutuhkan
5. Memasukkan berat ikan ke wadah dan menghitung total berat ikan per wadah
6. Menghitung rata rata berat ikan pada wadah
7. Tampilkan total berat ikan per wadah dan rata rata berat ikan per wadah

Cara kerja dari program ini adalah user menginputkan jumlah ikan dan jumlah ikan yang di dalam wadah. Program meminta user untuk menginputkan berat dari ikan-ikan. Setelah itu, program menghitung jumlah wadah yang dibutuhkan dan membagikan ikan-ikan itu ke dalam wadah. Lalu, menghitung total berat ikan per wadah. Di akhir program menghitung rata-rata ikan pada setiap wadah.

**3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.**

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / IF-11-05

package main

import "fmt"

// Tipe data array untuk menyimpan berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat min dan max
func hitungMinMax(arrBerat arrBalita, n int) (minBerat, maxBerat float64) {
	if n == 0 {
		return 0, 0 
	}

	// untuk nilai minimum dan maksimum
	minBerat, maxBerat = arrBerat[0], arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < minBerat {
			minBerat = arrBerat[i]
		}
		if arrBerat[i] > maxBerat {
			maxBerat = arrBerat[i]
		}
	}
	return
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arrBerat arrBalita, n int) float64 {
	if n == 0 {
		return 0
	}

	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var dataBerat arrBalita

	for i := 0; i < len(dataBerat); i++ {
		dataBerat[i] = 0
	}

	// Input jumlah balita yang ingin di data
	fmt.Print("Banyak data berat balita : ")
	fmt.Scan(&n)

	// mevalidasi jumlah data
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data tidak lebih dari 100.")
		return
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke - %d : ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	// untuk menghitung minimum, maksimum, dan rata-rata berat
	minBerat, maxBerat := hitungMinMax(dataBerat, n)
	rataRata := hitungRerata(dataBerat, n)

	// Output akhir
	fmt.Printf("\nBerat Balita Minimum : %.2f kg\n", minBerat)
	fmt.Printf("Berat Balita Maksimal : %.2f kg\n", maxBerat)
	fmt.Printf("Berat Balita Rata-Rata : %.2f kg\n", rataRata)
}
```
#### Screenshoot Source Code
![Unguided3 carbon](https://github.com/user-attachments/assets/b410880b-5435-49b8-953d-8a9b95a91899)

#### Screenshoot Output
![Unguided3 go](https://github.com/user-attachments/assets/b6314948-a84d-4c4b-be95-598954e40198)

#### Deskripsi Program
Program ini merupakan program untuk menghitung rata-rata dan menampilkan berat terkecil dan terbesar. Program menggunakan algoritma pencarian linear. Program juga menggunakan array bertipe data 'arrBalita' untuk menyimpan angka berat badan dengan tipe data 'float64'

Algoritma Program :
1. Input jumlah balita yang ingin diitung
2. Validasi jumlah balita tidak lebih dari 100
3. Inputkan berat balita pada setiap balita
4. Gunakan fungsi 'hitungMinMax()' untuk mencari berat minimal dan maksimum
5. Hitung rata rata dengan fungsi 'hitungRerata()'
6. Tampilkan hasil akhir yaitu berat balita minimum dan maksimum dan rata ratanya

Cara kerja dari program ini adalah user menginputkan jumlah balita yang ingin dihitung. user menginputkan berat untuk setiap balita. Program menghitung nilai maksimum, minimum, dan rata-rata berat balita.

## Referensi
[1] Pratheesh, P. C. (2024). Exploring Sorting and Searching Algorithms with Golang: Linear Search. Diakses dari https://pcpratheesh-medium-com.translate.goog/exploring-sorting-and-searching-algorithms-with-golang-linear-search-881e992e4e0c?_x_tr_sl=en&_x_tr_tl=id&_x_tr_hl=id&_x_tr_pto=sge&_x_tr_hist=true                                                        

[2] Pratheesh, P. C. (2024). Exploring Sorting and Searching Algorithms with Golang: Binary Search. Diakses dari https://pcpratheesh-medium-com.translate.goog/exploring-sorting-and-searching-algorithms-with-golang-binary-search-6679716bcd3a?_x_tr_sl=en&_x_tr_tl=id&_x_tr_hl=id&_x_tr_pto=sge&_x_tr_hist=true 
