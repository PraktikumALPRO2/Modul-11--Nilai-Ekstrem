<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XI</strong></h2>
<h2 align="center"><strong> PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA </strong></h2> 

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fadilah Salehah / 2311102164<br>
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
1. [Tujuan Praktikum](#tujuan-praktikum)
2. [Dasar Teori](#dasar-teori)
3. [Unguided](#unguided)
4. [Kesimpulan](#kesimpulan)
5. [Daftar Pustaka](#daftar-pustaka)

## Tujuan Praktikum
1. Mahasiswa menggunakan fungsi dalam modularisasi program agar program lebih terstruktur dan mudah dipahami.
2. Mahasiswa mengimplementasikan algoritma untuk mencari nilai minimum, maksimum, dan rata-rata dari sekumpulan data.

## Dasar Teori

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

  
**Contoh Implementasi nya program pencarian nilai ekstrem dapat dilakukan dengan iterasi sederhana:**
```go
package main

import "fmt"

func findMaxMin(arr []int) (int, int) {
    max := arr[0]
    min := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    return max, min
}

func main() {
    data := []int{10, 3, 5, 8, 2, 15, 7}
    max, min := findMaxMin(data)
    fmt.Printf("Nilai maksimum: %d, Nilai minimum: %d\n", max, min)
}
```

Aplikasi Pencarian Nilai Ekstrim Pencarian nilai ekstrem sering digunakan dalam:

1. Analisis Data: Untuk menentukan range data atau mendeteksi outlier dalam dataset.
2. Pengolahan Citra Digital: Menentukan nilai intensitas piksel maksimum dan minimum.
3. Algoritma Optimasi: Mencari solusi optimal dalam himpunan solusi yang ada.

Kesimpulan :

Pencarian nilai ekstrem adalah prosedur dasar namun penting dalam analisis data dan pemrograman. Pemahaman terhadap metode dan implementasi pencarian nilai ekstrem membantu dalam pengolahan data dan pengambilan keputusan yang lebih baik dalam berbagai bidang aplikasi.


## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. 

**Masukan** terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. 

**Keluaran** terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

### Source Code :
```go
// Fadilah Salehah 
// 2311102164 
// S1 IF 11 05 

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("   Aplikasi Analisis Berat Kelinci")
		fmt.Println(strings.Repeat("=", 40))

		// Input: Jumlah kelinci
		var jumlah int
		fmt.Print("\nMasukkan jumlah kelinci: ")
		fmt.Scanln(&jumlah)

		// Validasi: Jumlah kelinci harus lebih dari 0
		if jumlah <= 0 {
			fmt.Println("\n[!] Jumlah kelinci tidak boleh kurang dari 1. Silakan coba lagi.")
			continue
		}

		// Membuat slice untuk berat kelinci
		beratKelinci := make([]float64, jumlah)
		fmt.Println("\nMasukkan berat kelinci satu per satu:")
		for i := 0; i < jumlah; i++ {
			fmt.Printf("  - Berat kelinci ke-%d: ", i+1)
			fmt.Scanln(&beratKelinci[i])
		}

		// Mengurutkan berat kelinci secara menaik
		sort.Float64s(beratKelinci)

		// Output: Menampilkan berat terkecil dan terbesar
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("         Ringkasan Data Berat")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("🐇 Berat terkecil : %.2f kg\n", beratKelinci[0])
		fmt.Printf("🐇 Berat terbesar : %.2f kg\n", beratKelinci[jumlah-1])
		fmt.Println("\n📋 Daftar berat kelinci (urut):")
		for i, berat := range beratKelinci {
			fmt.Printf("  %2d. %.2f kg\n", i+1, berat)
		}
		fmt.Println(strings.Repeat("-", 40))

		// Menanyakan apakah program akan diulangi
		var ulangi string
		fmt.Print("\n🔄 Ingin mengolah data lagi? (ya/tidak): ")
		fmt.Scanln(&ulangi)

		// Keluar jika jawaban bukan "ya"
		if strings.ToLower(ulangi) != "ya" {
			fmt.Println("\n✨ Terima kasih telah menggunakan aplikasi ini. ✨")
			break
		}
	}
}

```

### Output:
![image](https://github.com/user-attachments/assets/103c1054-e306-41e6-9945-539edc7a9d6b)


### Deskripsi Program : 
Program ini digunakan untuk mengolah berat kelinci yang dimasukkan oleh pengguna. Pengguna diminta untuk memasukkan jumlah kelinci terlebih dahulu, lalu memasukkan berat masing-masing kelinci satu per satu. Setelah semua data berat dimasukkan, program mengurutkan berat dalam urutan menaik, kemudian menampilkan berat terkecil, terbesar, dan daftar berat yang terurut. Program juga memberikan opsi kepada pengguna untuk mengulangi proses pengolahan data jika diinginkan. Program akan berakhir jika pengguna memilih untuk tidak mengulangi.

### 2.Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. 

**Masukan** terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
### Source Code :
```go
// Fadilah Salehah
// 2311102164
// S1 IF 11 5

package main

import (
	"fmt"
)

func main() {
	var jumlahIkan, kapasitasWadah int

	// Input jumlah ikan yang dijual dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan yang dijual: ")
	fmt.Scan(&jumlahIkan)
	fmt.Print("Masukkan kapasitas wadah (jumlah ikan per wadah): ")
	fmt.Scan(&kapasitasWadah)

	// Validasi input
	if jumlahIkan <= 0 || kapasitasWadah <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari nol.")
		return
	}

	// Input berat masing-masing ikan
	beratIkan := make([]float64, jumlahIkan)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Pengelompokan ikan ke dalam wadah
	var totalBeratPerWadah []float64
	for i := 0; i < jumlahIkan; i += kapasitasWadah {
		batasAkhir := i + kapasitasWadah
		if batasAkhir > jumlahIkan {
			batasAkhir = jumlahIkan
		}

		// Hitung total berat ikan dalam satu wadah
		var total float64
		for j := i; j < batasAkhir; j++ {
			total += beratIkan[j]
		}
		totalBeratPerWadah = append(totalBeratPerWadah, total)
	}

	// Menampilkan hasil total berat dan rata-rata berat setiap wadah
	fmt.Println("\nHasil pengelompokan ikan:")
	for i, total := range totalBeratPerWadah {
		rataRata := total / float64(kapasitasWadah)
		if i == len(totalBeratPerWadah)-1 && jumlahIkan%kapasitasWadah != 0 {
			rataRata = total / float64(jumlahIkan%kapasitasWadah) // Penyesuaian untuk wadah terakhir
		}
		fmt.Printf("Wadah ke-%d: Total berat = %.2f kg, Rata-rata = %.2f kg\n", i+1, total, rataRata)
	}
}

```

### Output:
![image](https://github.com/user-attachments/assets/0a49c2a1-0846-4e10-bff9-fa15b70a5ee7)


### Deskripsi Program : 
Program ini digunakan untuk mengolah data berat ikan yang akan dijual dan membagi berat tersebut ke dalam sejumlah wadah. Pengguna diminta untuk memasukkan jumlah ikan dan berat masing-masing ikan, serta jumlah wadah yang akan digunakan. Program kemudian menghitung dan menampilkan total berat ikan di setiap wadah serta rata-rata berat ikan. Hasil ini membantu dalam mengetahui distribusi berat ikan yang lebih merata di setiap wadah.


### 3.Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya. Buatlah program dengan spesifikasi subprogram sebagai berikut:

![image](https://github.com/user-attachments/assets/d3e9f5b8-6afd-47e7-adad-ee2120145afe)

![image](https://github.com/user-attachments/assets/6fa77412-24ea-487a-8cda-315ecc20ef57)

![image](https://github.com/user-attachments/assets/77a11930-d333-40eb-827b-a524eebac368)

### Source Code :
```go
// Fadilah Salehah
// 2311102164
// S1 IF 11 5
package main

import (
	"fmt"
)

// Definisi tipe data untuk array berat balita
type arrayBerat [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func cariMinMax(berat arrayBerat, jumlah int, min *float64, max *float64) {
	*min = berat[0]
	*max = berat[0]

	for i := 1; i < jumlah; i++ {
		if berat[i] < *min {
			*min = berat[i]
		}
		if berat[i] > *max {
			*max = berat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRataRata(berat arrayBerat, jumlah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlah; i++ {
		total += berat[i]
	}
	return total / float64(jumlah)
}

// Fungsi utama
func main() {
	var jumlahBalita int
	var beratBalita arrayBerat

	// Input jumlah balita
	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahBalita)

	// Validasi jumlah balita
	if jumlahBalita <= 0 || jumlahBalita > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	fmt.Println("Masukkan berat balita:")
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Deklarasi variabel untuk hasil
	var beratMin, beratMax float64

	// Panggil fungsi cariMinMax
	cariMinMax(beratBalita, jumlahBalita, &beratMin, &beratMax)

	// Panggil fungsi hitungRataRata
	rataRata := hitungRataRata(beratBalita, jumlahBalita)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}


```

### Output:
![image](https://github.com/user-attachments/assets/0bc15fb0-8d70-4165-b41d-6825a96e6579)


### Deskripsi Program : 
Program ini dirancang untuk mengolah data berat balita yang dimasukkan oleh pengguna. Pengguna diminta memasukkan jumlah data berat balita dan berat masing-masing balita satu per satu. Program kemudian menghitung dan menampilkan berat minimum, maksimum, dan rata-rata dari data yang dimasukkan. Program ini berguna untuk mengetahui sebaran dan ringkasan statistik dari data berat balita.

## Kesimpulan 
Berdasarkan hasil praktikum, dapat disimpulkan bahwa:

1. Praktikum ini memberikan pemahaman tentang penggunaan array dalam menyimpan dan mengelola data numerik dalam program. Array memudahkan pengelolaan sejumlah data yang memiliki tipe yang sama dalam satu wadah, serta memungkinkan pengolahan data secara lebih efisien.

2. Selain itu, praktikum ini mengajarkan pentingnya penggunaan fungsi untuk memecah tugas-tugas tertentu dalam program, seperti menghitung nilai minimum, maksimum, dan rata-rata, yang membuat kode lebih terstruktur dan mudah dipahami.

3. Dalam praktikum, kita juga mempelajari cara mengelola input dari pengguna untuk mengisi array dan pentingnya menangani input yang tidak valid dengan pemeriksaan error agar program dapat berjalan dengan baik.

## Daftar Pustaka

[1] A. A. A. Donovan and B. W. Kernighan, The Go Programming Language. Boston, MA: Addison-Wesley, 2015.
