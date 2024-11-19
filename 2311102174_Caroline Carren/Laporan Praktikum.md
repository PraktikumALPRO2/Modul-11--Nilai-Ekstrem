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
  Caroline Carren A.R / 2311102174<br>
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
1. Mahasiswa mempelajari cara mengelola dan memproses data numerik menggunakan array.
2. Mahasiswa mengimplementasikan algoritma untuk mencari nilai minimum, maksimum, dan rata-rata dari sekumpulan data.
3. Mahasiswa menggunakan fungsi dalam modularisasi program agar program lebih terstruktur dan mudah dipahami.

## Dasar Teori

**Pengertian Nilai Ekstrim**

Nilai ekstrem dalam konteks himpunan data merujuk pada elemen data yang memiliki nilai maksimum atau minimum. 
Nilai maksimum adalah elemen terbesar dalam himpunan data, sedangkan nilai minimum adalah elemen terkecil. 
Pencarian nilai ekstrem sangat penting dalam berbagai bidang seperti statistik, sains data, dan pemrograman karena memberikan informasi mengenai batasan dan distribusi data.

Metode Pencarian Nilai Ekstrim Ada berbagai metode yang digunakan untuk menemukan nilai ekstrem dalam himpunan data. Metode ini dapat diterapkan secara manual melalui iterasi atau menggunakan fungsi-fungsi bawaan dalam bahasa pemrograman. 

- Pencarian Linier: Metode paling sederhana di mana setiap elemen dalam himpunan data diperiksa satu per satu untuk membandingkan nilainya. Kompleksitas waktu metode ini adalah O(n), di mana n adalah jumlah elemen dalam himpunan data.
- Pencarian dengan Algoritma Tertentu: Dalam dataset besar, algoritma yang lebih kompleks seperti algoritma divide and conquer dapat digunakan untuk mengoptimalkan pencarian nilai ekstrem dengan membagi data menjadi bagian-bagian yang lebih kecil.
  
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
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("    Program Pengolahan Berat Kelinci")
		fmt.Println(strings.Repeat("=", 40))

		// Input: Jumlah kelinci
		var n int
		fmt.Print("\nMasukkan jumlah kelinci: ")
		fmt.Scanln(&n)

		// Memastikan jumlah kelinci lebih dari 0
		if n <= 0 {
			fmt.Println("\n[!] Jumlah kelinci harus lebih dari 0. Silakan coba lagi.")
			continue
		}

		// Membuat slice untuk menyimpan berat kelinci
		weights := make([]float64, n)
		fmt.Println("\nMasukkan berat masing-masing kelinci:")
		for i := 0; i < n; i++ {
			fmt.Printf("  - Berat kelinci %d: ", i+1)
			fmt.Scanln(&weights[i])
		}

		// Mengurutkan berat kelinci dalam urutan menaik
		sort.Float64s(weights)

		// Output: Berat kelinci terkecil dan terbesar
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("        Hasil Pengolahan")
		fmt.Println(strings.Repeat("-", 40))
		fmt.Printf("🐇 Berat terkecil : %.2f kg\n", weights[0])
		fmt.Printf("🐇 Berat terbesar : %.2f kg\n", weights[n-1])
		fmt.Println("\n📋 Daftar berat kelinci (terurut):")
		for i, weight := range weights {
			fmt.Printf("  %2d. %.2f kg\n", i+1, weight)
		}
		fmt.Println(strings.Repeat("-", 40))

		// Menanyakan apakah pengguna ingin mengulangi
		var again string
		fmt.Print("\n🔄 Apakah Anda ingin memasukkan data lagi? (ya/tidak): ")
		fmt.Scanln(&again)

		// Jika pengguna tidak ingin mengulangi, keluar dari loop
		if strings.ToLower(again) != "ya" {
			fmt.Println("\n✨ Terima kasih telah menggunakan program ini. ✨")
			break
		}
	}
}
```

### Output:
![image](https://github.com/user-attachments/assets/770a4798-7034-4fe5-a61f-f3b40f5d3afd)

### Full code Screenshot :
![carbon](https://github.com/user-attachments/assets/e936c3cf-069c-4f1a-af22-ddf4502fb739)

### Deskripsi Program : 
Program ini digunakan untuk mengolah berat kelinci yang dimasukkan oleh pengguna. Pengguna diminta untuk memasukkan jumlah kelinci terlebih dahulu, lalu memasukkan berat masing-masing kelinci satu per satu. Setelah semua data berat dimasukkan, program mengurutkan berat dalam urutan menaik, kemudian menampilkan berat terkecil, terbesar, dan daftar berat yang terurut. Program juga memberikan opsi kepada pengguna untuk mengulangi proses pengolahan data jika diinginkan. Program akan berakhir jika pengguna memilih untuk tidak mengulangi.

### Algoritma Program :
- Deklarasi Variabel:
  1. n adalah variabel integer yang menyimpan jumlah kelinci.
  2. weights adalah slice bertipe float64 yang digunakan untuk menyimpan berat kelinci yang dimasukkan pengguna.
  3. again adalah variabel string untuk menyimpan pilihan pengguna terkait pengulangan program.

- Fungsi Utama (main):
  1. Cetak judul program.
  2. Minta input jumlah kelinci (n)
     - Jika n ≤ 0, tampilkan pesan kesalahan dan ulangi proses input.
  3. Buat slice weights untuk menyimpan berat kelinci.
  4. Minta input berat kelinci satu per satu dan simpan di weights.
  5. Urutkan slice weights dalam urutan menaik menggunakan sort.Float64s().
  6. Tampilkan hasil:
     - Berat terkecil (weights[0]).
     - Berat terbesar (weights[n-1]).
     - Daftar berat kelinci yang terurut.
  7. Tanyakan apakah pengguna ingin mengulangi proses.
     - Jika jawaban "ya", kembali ke langkah 2.
     - Jika jawaban "tidak", cetak pesan terima kasih dan keluar dari program.

### Cara Kerja Program :
1. Input Jumlah Kelinci:
   - Program meminta pengguna untuk memasukkan jumlah kelinci yang akan diproses. Jika input tidak valid (≤ 0), program meminta input ulang.
     
2. Input Berat Kelinci:
   - Program meminta pengguna memasukkan berat tiap kelinci satu per satu hingga jumlah yang dimasukkan.

3. Mengurutkan dan Menampilkan Hasil:
   - Berat kelinci diurutkan dalam urutan menaik menggunakan sort.Float64s().
   - Program menampilkan berat terkecil, terbesar, dan daftar berat yang terurut.

4. Memeriksa Ulangi Proses:
   - Program bertanya kepada pengguna apakah ingin mengulangi proses pengolahan data. Jika "ya", program mengulangi; jika "tidak", program mengakhiri proses dengan menampilkan pesan terima kasih.

5. Output Hasil:
   - Program mencetak berat terkecil, terbesar, dan daftar berat kelinci yang terurut, serta menampilkan apakah pengguna ingin mengulangi proses atau tidak.
     
### 2.Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. 

**Masukan** terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

### Source Code :
```go
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input: Masukan
	var x, y int
	var ikan []float64

	// Meminta pengguna untuk memasukkan jumlah ikan
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("    Program Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("\nMasukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scanln(&x)
	if x <= 0 {
		fmt.Println("\n[!] Jumlah ikan harus positif.")
		return // Keluar dari program jika jumlah ikan tidak valid
	}

	// Meminta pengguna untuk memasukkan jumlah wadah
	fmt.Print("\nMasukkan jumlah wadah (y): ")
	fmt.Scanln(&y)
	if y <= 0 {
		fmt.Println("\n[!] Jumlah wadah harus positif.")
		return // Keluar dari program jika jumlah wadah tidak valid
	}

	// Meminta pengguna untuk memasukkan berat setiap ikan satu per satu
	fmt.Println("\nMasukkan berat setiap ikan (dalam kg):")
	for i := 0; i < x; i++ {
		var berat float64
		fmt.Printf("  - Berat ikan %d: ", i+1)
		_, err := fmt.Scanln(&berat)
		if err != nil {
			fmt.Println("\n[!] Input berat ikan tidak valid.")
			return // Keluar dari program jika ada kesalahan input
		}
		ikan = append(ikan, berat)
	}

	// Hitung total berat ikan di setiap wadah
	wadah := make([]float64, y) // Membuat slice untuk menyimpan total berat di setiap wadah
	for i := 0; i < x; i++ {
		wadah[i%y] += ikan[i] // Menambahkan berat ikan ke wadah yang sesuai
	}

	// Hitung rata-rata berat ikan di setiap wadah
	rataRata := 0.0
	for _, v := range wadah {
		rataRata += v // Menjumlahkan total berat di semua wadah
	}
	rataRata /= float64(x) // Rata-rata berdasarkan jumlah ikan

	// Output: Keluaran
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("     Hasil Pengolahan Berat Ikan")
	fmt.Println(strings.Repeat("-", 40))

	// Menampilkan total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, v := range wadah {
		fmt.Printf("  Wadah %d: %.2f kg\n", i+1, v) // Menampilkan total berat ikan di setiap wadah
	}

	// Menampilkan berat rata-rata ikan di setiap wadah
	fmt.Println("\nBerat rata-rata ikan di setiap wadah:")
	fmt.Printf("  %.2f kg\n", rataRata) // Menampilkan rata-rata berat ikan

	// Akhir dari program
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Terima kasih telah menggunakan program ini!")
	fmt.Println(strings.Repeat("=", 40))
}
```

### Output:
![image](https://github.com/user-attachments/assets/97bf88b8-658d-4685-8837-d167d925b0c3)

### Full code Screenshot :
![carbon (2)](https://github.com/user-attachments/assets/20c9af56-91c5-4be4-8586-24d765c93691)

### Deskripsi Program : 
Program ini digunakan untuk mengolah data berat ikan yang akan dijual dan membagi berat tersebut ke dalam sejumlah wadah. Pengguna diminta untuk memasukkan jumlah ikan dan berat masing-masing ikan, serta jumlah wadah yang akan digunakan. Program kemudian menghitung dan menampilkan total berat ikan di setiap wadah serta rata-rata berat ikan. Hasil ini membantu dalam mengetahui distribusi berat ikan yang lebih merata di setiap wadah.

### Algoritma Program :
- Deklarasi Variabel:
  1. x adalah variabel integer untuk menyimpan jumlah ikan.
  2. y adalah variabel integer untuk menyimpan jumlah wadah.
  3. ikan adalah slice bertipe float64 yang menyimpan berat masing-masing ikan.
  4. wadah adalah slice bertipe float64 untuk menyimpan total berat ikan di setiap wadah.
  5. rataRata adalah variabel float64 untuk menghitung rata-rata berat ikan.

- Langkah-langkah Program:
  1. Input Data:
     - Minta pengguna memasukkan jumlah ikan (x). Jika x ≤ 0, tampilkan pesan kesalahan dan keluar dari program.
     - Minta pengguna memasukkan jumlah wadah (y). Jika y ≤ 0, tampilkan pesan kesalahan dan keluar dari program.

2. Input Berat Ikan:
   - Minta pengguna untuk memasukkan berat setiap ikan satu per satu dan simpan dalam slice ikan. Jika input tidak valid, program menampilkan pesan kesalahan dan keluar.

3. Pengolahan Data:
   - Inisialisasi slice wadah untuk menyimpan total berat di setiap wadah.
   - Distribusikan berat ikan ke dalam wadah secara berurutan menggunakan indeks modulo (i % y).

4. Menghitung Rata-rata:
   - Jumlahkan total berat ikan di semua wadah.
   - Hitung rata-rata berat ikan berdasarkan jumlah ikan.

5. Output Data:
   - Tampilkan total berat ikan di setiap wadah.
   - Tampilkan rata-rata berat ikan di setiap wadah.

6. Akhir Program:
   - Tampilkan pesan akhir bahwa program telah selesai.

### Cara Kerja Program :
1. Input Jumlah Ikan dan Wadah:
   - Pengguna diminta untuk memasukkan jumlah ikan (x) dan jumlah wadah (y). Jika input tidak valid, program keluar dengan pesan kesalahan.

2. Input Berat Ikan:
   - Pengguna memasukkan berat ikan satu per satu, yang disimpan dalam slice ikan.

3. Distribusi Berat ke Wadah:
   - Berat ikan didistribusikan ke dalam wadah berdasarkan indeks modulo (i % y) untuk memastikan pembagian yang merata.

4. Penghitungan Rata-rata:
   - Total berat di semua wadah dijumlahkan, kemudian dihitung rata-rata berat ikan.

5. Output Hasil:
   - Program menampilkan total berat ikan di setiap wadah dan rata-rata berat ikan secara keseluruhan.

6. Tampilan Akhir:
   - Program menampilkan pesan terima kasih dan selesai.

### 3.Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya. Buatlah program dengan spesifikasi subprogram sebagai berikut:

![image](https://github.com/user-attachments/assets/d3e9f5b8-6afd-47e7-adad-ee2120145afe)

![image](https://github.com/user-attachments/assets/6fa77412-24ea-487a-8cda-315ecc20ef57)

![image](https://github.com/user-attachments/assets/77a11930-d333-40eb-827b-a524eebac368)

### Source Code :
```go
// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"strings"
)

// Definisi tipe data untuk array berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
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

// Fungsi untuk menghitung rata-rata berat balita
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
	var bMin, bMax float64

	// Tampilan header
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("  Program Pengolahan Data Berat Balita")
	fmt.Println(strings.Repeat("=", 40))

	// Meminta pengguna untuk memasukkan jumlah data balita
	fmt.Print("\nMasukan banyak data berat balita: ")
	fmt.Scanln(&n)

	// Memasukkan data berat balita satu per satu
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}

	// Memanggil fungsi untuk menghitung minimum dan maksimum
	hitungMinMax(berat, n, &bMin, &bMax)

	// Menghitung rata-rata berat balita
	rata := rataRata(berat, n)

	// Menampilkan hasil
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("   Hasil Pengolahan Berat Balita")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("🌟 Berat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("🌟 Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("🌟 Rerata berat balita   : %.2f kg\n", rata)
	fmt.Println(strings.Repeat("-", 40))

	// Menutup program dengan ucapan terima kasih
	fmt.Println("\nTerima kasih telah menggunakan program ini!")
	fmt.Println(strings.Repeat("=", 40))
}
```

### Output:
![image](https://github.com/user-attachments/assets/81f4450f-a0c3-4186-945f-ddccffcc84e6)

### Full code Screenshot :
![carbon (3)](https://github.com/user-attachments/assets/b92c399d-d911-4db9-8d89-f3e5927a3575)

### Deskripsi Program : 
Program ini dirancang untuk mengolah data berat balita yang dimasukkan oleh pengguna. Pengguna diminta memasukkan jumlah data berat balita dan berat masing-masing balita satu per satu. Program kemudian menghitung dan menampilkan berat minimum, maksimum, dan rata-rata dari data yang dimasukkan. Program ini berguna untuk mengetahui sebaran dan ringkasan statistik dari data berat balita.

### Algoritma Program :
- Deklarasi Variabel:
  1. n adalah variabel integer yang menyimpan jumlah balita.
  2. berat adalah array bertipe arrBalita untuk menyimpan berat balita hingga maksimal 100 data.
  3. bMin dan bMax adalah variabel float64 untuk menyimpan berat minimum dan maksimum balita.

- Fungsi hitungMinMax:
  1. Inisialisasi bMin dan bMax dengan nilai elemen pertama dari array arrBerat.
  2. Iterasi mulai dari elemen kedua hingga elemen terakhir.
  3. Jika elemen saat ini lebih kecil dari bMin, perbarui bMin.
  4. Jika elemen saat ini lebih besar dari bMax, perbarui bMax.

- Fungsi rataRata:
  1. Inisialisasi total dengan 0.
  2. Lakukan iterasi untuk menjumlahkan semua elemen dalam array arrBerat.
  3. Hitung dan kembalikan hasil rata-rata dengan membagi total dengan jumlah elemen n.

- Program Utama (main):
  1. Cetak header program.
  2. Minta pengguna memasukkan jumlah data berat balita (n).
  3. Minta pengguna untuk memasukkan berat masing-masing balita satu per satu dan simpan di array berat.
  4. Panggil fungsi hitungMinMax untuk menghitung berat minimum dan maksimum.
  5. Panggil fungsi rataRata untuk menghitung rata-rata berat.
  6. Cetak hasil perhitungan berat minimum, maksimum, dan rata-rata.
  7. Tampilkan ucapan terima kasih dan akhiri program.

### Cara Kerja Program :
1. Input Data Berat Balita:
   - Pengguna diminta untuk memasukkan jumlah balita (n) dan kemudian berat masing-masing balita satu per satu.

2. Menghitung Berat Minimum dan Maksimum:
   - Program memanggil fungsi hitungMinMax untuk mencari berat balita terkecil dan terbesar dalam array.

3. Menghitung Rata-rata Berat:
   - Program memanggil fungsi rataRata untuk menghitung rata-rata berat balita.

4. Output Hasil:
   - Program menampilkan hasil berupa berat minimum, maksimum, dan rata-rata.

5. Penutup:
   - Program mencetak pesan terima kasih sebagai tanda akhir.

## Kesimpulan 
Berdasarkan hasil praktikum, dapat disimpulkan bahwa:

1. Praktikum ini memberikan pemahaman tentang penggunaan array dalam menyimpan dan mengelola data numerik dalam program. Array memudahkan pengelolaan sejumlah data yang memiliki tipe yang sama dalam satu wadah, serta memungkinkan pengolahan data secara lebih efisien.

2. Selain itu, praktikum ini mengajarkan pentingnya penggunaan fungsi untuk memecah tugas-tugas tertentu dalam program, seperti menghitung nilai minimum, maksimum, dan rata-rata, yang membuat kode lebih terstruktur dan mudah dipahami.

3. Dalam praktikum, kita juga mempelajari cara mengelola input dari pengguna untuk mengisi array dan pentingnya menangani input yang tidak valid dengan pemeriksaan error agar program dapat berjalan dengan baik.

## Daftar Pustaka

[1] A. A. A. Donovan and B. W. Kernighan, The Go Programming Language. Boston, MA: Addison-Wesley, 2015.
