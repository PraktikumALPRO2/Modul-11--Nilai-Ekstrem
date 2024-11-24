<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 11</h2>
<h2 align="center">NILAI EKSTREM</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="250"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI
####  Ide Pencarian Nilai Max/Min
Pencarian adalah suatu proses yang biasa dilakukan setiap hari. Contohnya dalam dunia nyata sangat beragam. Salah satu contohnya adalah mencari file di dalam direktori komputer, mencari teks dalam dokumen, mencari buku di rak buku, dan sebagainya. Salah satu algoritma pencarian nilai terkecil atau terbesar pada sekumpulan data juga dikenal sebagai pencarian nilai ekstrim.

Ide algoritma sangat sederhana. Nilai atau indeks ke nilai tertinggi dari data yang telah diproses disimpan untuk dibandingkan dengan data berikutnya karena data diproses secara sekuensial. Nilai yang berhasil disimpan sampai algoritma berakhir, yang merupakan nilai tertinggi yang dicari. Adapun algoritmanya secara umum adalah sebagai berikut:
1. Jadikan data pertama sebagai nilai ekstrim
2. Lakukan validasi nilai ekstrim dari data kedua hingga data terakhir.
Apabila nilai ekstrim tidak valid, maka update nilai ekstrims tersebut dengan data data yang dicek.
4. Apabila semua data telah dicek, maka nilai ekstrim yang dimiliki adalah valid motics lab.

Berikut ini adalah notasi dalam pseudocode dan bahasa Go, misalnya untuk pencarian nilai terbesar atau maksimum:
![Screenshot 2024-11-24 230114](https://github.com/user-attachments/assets/8fa378ae-4a18-402d-8474-bc06affde444)


<br></br>


#### II. UNGUIDED

##### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
```go
package main

import "fmt"

func main() {

	var berat [1000]float64

	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	minBerat := berat[0]
	maxBerat := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}

```
##### Screenshoot Output
![Screenshot 2024-11-24 220950](https://github.com/user-attachments/assets/85ee6b9a-1ac0-4e05-97ef-b4435ad3ac01)

##### Deskripsi Program
Program ini digunakan untuk mencatat dan menganalisis berat anak kelinci. User memasukkan jumlah anak kelinci, kemudian berat masing-masing anak kelinci dimasukkan secara berurutan. Data berat anak kelinci disimpan dalam sebuah array. Selanjutnya, program mencari dan menentukan berat terkecil dan terbesar dari data yang dimasukkan. Akhirnya, program mencetak berat terkecil dan terbesar.

##### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. informatic! 100 Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
```go
package main

import "fmt"

func main() {

	var x, y int
	var totalBeratWadah []float64
	var rataRataWadah []float64

	fmt.Print("Masukkan x dan y (jumlah ikan dan kapasitas wadah): ")
	fmt.Scanf("%d %d", &x, &y)

	beratIkan := make([]float64, x)
	fmt.Print("Masukkan berat ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scanf("%f", &beratIkan[i])
	}

	for i := 0; i < x; i += y {
		end := i + y
		if end > x {
			end = x
		}

		wadah := beratIkan[i:end]
		totalBerat := 0.0

		for _, berat := range wadah {
			totalBerat += berat
		}

		totalBeratWadah = append(totalBeratWadah, totalBerat)

		if len(wadah) > 0 {
			rataRata := totalBerat / float64(len(wadah))
			rataRataWadah = append(rataRataWadah, rataRata)
		} else {
			rataRataWadah = append(rataRataWadah, 0)
		}
	}

	fmt.Println("Total berat ikan di setiap wadah:", totalBeratWadah)
	fmt.Println("Rata-rata berat ikan di setiap wadah:", rataRataWadah)
}

```
##### Screenshoot Output
![Screenshot 2024-11-24 225057](https://github.com/user-attachments/assets/9ea4babf-42fa-4ddb-a0b3-dfaf1f194aad)

##### Deskripsi Program
Program ini bertujuan untuk menghitung total berat ikan dan rata-rata berat ikan di setiap wadah berdasarkan jumlah ikan dan kapasitas wadah yang ditentukan. Program dimulai dengan meminta input jumlah ikan (x) dan kapasitas setiap wadah (y), diikuti oleh data berat masing-masing ikan. Berat ikan dimasukkan ke dalam array dan dikelompokkan ke dalam wadah sesuai kapasitas y. Untuk setiap wadah, program menghitung total berat ikan dan rata-rata berat per wadah, kemudian menyimpan hasilnya dalam dua array terpisah. Hasil akhir berupa daftar total berat dan rata-rata berat ikan di setiap wadah.

##### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {

	var n int
	var beratBalita arrBalita
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	hitungMinMax(beratBalita, n, &bMin, &bMax)
	rata := rerata(beratBalita, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}

```
##### Screenshoot Output
![Screenshot 2024-11-24 221053](https://github.com/user-attachments/assets/e24c449a-fa1d-4a68-bb08-aa547968a903)

##### Deskripsi Program
Program ini digunakan untuk mencatat dan menganalisis data berat badan balita. Program menerima input berupa jumlah balita dan berat masing-masing balita, lalu menghitung berat balita terkecil, terbesar, dan rata-rata berat balita. Data berat disimpan dalam array. Fungsi hitungMinMax adalah menghitung berat terkecil dan terbesar, sedangkan fungsi rerata menghitung nilai rata-rata berat balita berdasarkan total berat yang dibagi jumlah data. Hasil analisis berupa berat minimum, maksimum, dan rata-rata.


### Referensi
[1] MODUL 11 NILAI EKSTREM

