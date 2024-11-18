<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL VII</strong></h2>
<h2 align="center"><strong> STRUCT DAN ARRAY </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Wahyu Hidayat / 2311102178<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
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
### Definisi Tipe Bentukan
Tipe bentukan adalah fitur dalam bahasa pemrograman yang memungkinkan pengguna untuk mendefinisikan tipe data baru yang dapat menyesuaikan kebutuhan program. Tipe bentukan ini terbagi menjadi dua jenis utama, yaitu alias dan struct [1].

#### 1. Alias (Type): 
Alias adalah teknik yang digunakan untuk memberi nama baru yang lebih ringkas atau familiar pada suatu tipe data yang sudah ada. Contohnya, integer bisa diberi alias sebagai bilangan. Biasanya, penggantian nama ini dilakukan dengan menggunakan kata kunci type, yang mempermudah pengkodean dan meningkatkan keterbacaan [2].

#### 2. Struct atau Record: 
Struct adalah tipe data yang mengelompokkan beberapa variabel dengan berbagai tipe data menjadi satu kesatuan. Masing-masing data dalam struct tersimpan dalam field yang berbeda, yang berguna untuk menyimpan data yang memiliki hubungan erat, seperti data mahasiswa. Pada beberapa bahasa pemrograman, kesamaan tipe dari dua variabel struct ditentukan oleh struktur dari field-nya, bukan nama struct-nya [3].

### Array Statis dan Dinamis
Array adalah struktur data yang digunakan untuk menyimpan data dengan tipe yang sama dalam jumlah yang tetap atau berubah-ubah. Terdapat dua jenis array dalam pemrograman:

#### 1. Array Statis: 
Array statis adalah array dengan ukuran tetap yang ditentukan saat deklarasi. Setiap elemen array dapat diakses menggunakan indeks yang dimulai dari 0, misalnya int nilai[5] = {85, 90, 78, 88, 92}; [4].

#### 2. Slice (Array Dinamik): 
Dalam bahasa pemrograman seperti Go, terdapat array dinamis atau slice, yang memungkinkan ukuran array untuk berubah-ubah selama runtime. Deklarasi slice mirip dengan array, namun tanpa jumlah elemen tertentu. Slice dapat dialokasikan dengan fungsi make, sementara fungsi len dan cap dapat digunakan untuk mengetahui jumlah elemen dan kapasitas slice. Fungsi append dapat digunakan untuk menambah elemen baru ke dalam slice, menyesuaikan ukuran array dinamis tersebut sesuai kebutuhan [5].

### Map (Array Dinamik dengan Kunci)
Map adalah tipe data yang memungkinkan penggunaan kunci (key) selain integer sebagai indeks. Kunci dalam map dapat berupa tipe data apa saja, seperti string atau integer, yang membuatnya sangat fleksibel untuk menyimpan data dengan pasangan kunci-nilai (key-value) [6].

### Menggabungkan Struct dan Array
Struct dan array sering digunakan bersamaan untuk mengelola data yang kompleks. Contohnya, kita bisa menggunakan array berisi struct untuk menyimpan data mahasiswa dalam jumlah banyak. Contoh deklarasinya adalah:

#### Source Code
```go
struct Mahasiswa kelas[30];
```
Penggunaan ini memungkinkan kita untuk mengelola data dalam format yang terstruktur dan teratur, serta lebih mudah diakses dalam berbagai skenario pemrograman [7].

## II. GUIDED
## 1. Menghitung Durasi Parkir Berdasarkan Waktu Kedatangan dan Pulang
#### Source Code
```go
package main 
import "fmt" 
type waktu struct { 
    jam, menit, detik int 
} 
 
func main(){ 
    var wParkir, wPulang, durasi waktu 
    var dParkir, dPulang, lParkir int 
    fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik) 
    fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik) 
    dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600 
    dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600 
    lParkir = dPulang - dParkir 
    durasi.jam = lParkir / 3600 
    durasi.menit = lParkir % 3600 / 60 
    durasi.detik = lParkir % 3600 % 60 
    fmt.Printf("Lama parkir: %d jam %d menit %d detik", 
        durasi.jam, durasi.menit, durasi.detik) 
}  
```
#### Screenshoot Source Code
![Screenshot 2024-11-17 064624](https://github.com/user-attachments/assets/6db2201b-0817-4e8c-b1c5-a986171ee17e)


#### Screenshoot Output
![Screenshot 2024-11-17 064639](https://github.com/user-attachments/assets/e383294d-6cc8-407f-bc67-9602db42c479)

#### Deskripsi Program
Program ini digunakan untuk menghitung durasi waktu parkir berdasarkan waktu kedatangan dan waktu pulang yang diberikan dalam format jam, menit, dan detik. Hasil durasi ditampilkan dalam satuan jam, menit, dan detik.

#### Algoritma Program
1. Buat struktur waktu dengan atribut jam, menit, dan detik.
2. Baca input waktu parkir dan waktu pulang.
3. Konversi waktu parkir dan waktu pulang ke total detik.
4. Hitung selisih total detik untuk mendapatkan lama parkir.
5. Konversi selisih detik ke format jam, menit, dan detik.
6. Tampilkan hasil dalam format lama parkir.

#### Cara Kerja
1. Input waktu kedatangan dan pulang dalam format jam, menit, dan detik.
2. Waktu dikonversi ke detik total.
3. Selisih detik dihitung sebagai lama parkir.
4. Hasil konversi selisih detik ditampilkan dalam jam, menit, dan detik.

## 2. Menambahkan Nama ke Dalam Daftar Teman Tanpa Duplikasi
#### Source Code
```go
//Guided 2 - Slice
package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah nama sudah ada di dalam slice
func sudahAda(daftarTeman []string, nama string) bool {
	for _, teman := range daftarTeman {
		if teman == nama {
			return true
		}
	}
	return false
}

func main() {
	// Slice awal untuk daftar teman dengan beberapa data
	daftarTeman := []string{"Andi", "Budi", "Cici"}

	// Nama-nama baru yang ingin ditambahkan
	namaBaru := []string{"Dewi", "Budi", "Eka"}

	// Menambahkan nama baru hanya jika belum ada di daftar
	for _, nama := range namaBaru {
		if !sudahAda(daftarTeman, nama) {
			daftarTeman = append(daftarTeman, nama)
		} else {
			fmt.Println("Nama", nama, "sudah ada dalam daftar.")
		}
	}

	// Menampilkan daftar teman akhir
	fmt.Println("Daftar Teman:", daftarTeman)
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-17 073025](https://github.com/user-attachments/assets/b7fdc65b-a535-43fe-86d2-8b40050de2b9)

#### Screenshoot Output
![Screenshot 2024-11-17 073034](https://github.com/user-attachments/assets/fc6e1a63-6609-4af4-a1e9-522f49e919c6)

#### Deskripsi Program
Program ini memanipulasi slice untuk menyimpan daftar nama teman. Nama baru akan ditambahkan ke dalam daftar hanya jika nama tersebut belum ada, sehingga menghindari duplikasi. Jika nama sudah ada, program akan memberikan notifikasi kepada pengguna.

#### Algoritma Program
1. Buat slice awal yang berisi daftar nama teman.
2. Definisikan fungsi sudahAda untuk memeriksa apakah nama tertentu sudah ada di dalam slice.
3. Buat slice baru yang berisi nama-nama tambahan.
4. Periksa setiap nama dalam slice baru:
   - Jika belum ada dalam daftar teman, tambahkan ke slice.
   - Jika sudah ada, tampilkan pesan bahwa nama tersebut sudah ada.
5. Tampilkan daftar teman yang sudah diperbarui.

#### Cara Kerja
1. Program memulai dengan slice awal berisi nama-nama seperti "Andi", "Budi", dan "Cici".
2. Fungsi sudahAda digunakan untuk memeriksa keberadaan nama di dalam daftar.
3. Nama-nama baru seperti "Dewi", "Budi", dan "Eka" diperiksa:
   - "Dewi" dan "Eka" ditambahkan karena belum ada.
   - "Budi" tidak ditambahkan karena sudah ada, dan program menampilkan pesan.
4. Slice akhir, yang sudah diperbarui tanpa duplikasi, ditampilkan sebagai output.

## 3. Menampilkan Daftar Harga Buah Menggunakan Map

#### Source Code
```go
//Guided 2 - Map

package main

import (
	"fmt"
)

func main() {
	// Membuat map dengan nama buah sebagai kunci dan harga sebagai nilai
	hargaBuah := map[string]int{
		"Apel":  5000,
		"Pisang": 3000,
		"Mangga": 7000,
	}

	// Menampilkan harga dari setiap buah
	fmt.Println("Harga Buah:")
	for buah, harga := range hargaBuah {
		fmt.Printf("%s: Rp%d\n", buah, harga)
	}

	fmt.Print("Harga buah Mangga = ", hargaBuah["Mangga"])
}
```

#### Screenshoot Source Code
![Screenshot 2024-11-17 073501](https://github.com/user-attachments/assets/98e325a5-417b-4831-b9e7-e61c9bbbed99)

#### Screenshoot Output
![Screenshot 2024-11-17 073511](https://github.com/user-attachments/assets/1c009fe8-0cc0-4c67-9939-b77281ec16fc)

#### Deskripsi Program
Program ini menggunakan struktur data map di Go untuk menyimpan data berupa nama buah sebagai kunci dan harga sebagai nilai. Program menampilkan daftar harga semua buah dan mencetak harga spesifik untuk buah tertentu.

#### Algoritma Program
1. Buat map hargaBuah dengan kunci berupa nama buah dan nilai berupa harga buah.
2. Tambahkan beberapa data buah beserta harga ke dalam map.
3. Iterasi melalui map untuk mencetak semua buah beserta harganya.
4. Ambil dan tampilkan harga dari buah tertentu dengan menggunakan nama buah sebagai kunci.

#### Cara Kerja
1. Map hargaBuah dideklarasikan dengan beberapa data awal:
   - Apel berharga lima ribu.
   - Pisang berharga tiga ribu.
   - Mangga berharga tujuh ribu.
2. Program menggunakan perulangan for range untuk mencetak semua nama buah dan harga yang ada dalam map.
3. Harga dari buah tertentu, seperti Mangga, diakses langsung menggunakan kunci hargaBuah["Mangga"] dan ditampilkan.

## III. UNGUIDED
## 1. Pengecekan Posisi Titik terhadap Dua Lingkaran
#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

// Tipe data untuk menyimpan koordinat titik
type Titik struct {
	x, y int
}

// Tipe data untuk menyimpan lingkaran
type Lingkaran struct {
	pusat Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	// Input lingkaran 1
	var cx1, cy1, r1 int
	fmt.Printf("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1):")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input lingkaran 2
	var cx2, cy2, r2 int
	fmt.Printf("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik sembarang
	var px, py int
	fmt.Printf("Masukkan koordinat titik sembarang (px py):")
	fmt.Scan(&px, &py)

	// Buat lingkaran dan titik
	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{px, py}

	// Cek posisi titik terhadap lingkaran
	dalamL1 := didalam(lingkaran1, titik)
	dalamL2 := didalam(lingkaran2, titik)

	// Output posisi titik
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
#### Screenshoot Source Code
![Screenshot 2024-11-17 074851](https://github.com/user-attachments/assets/f0225f5b-fcc6-42c0-974b-add5176b83b3)

#### Screenshoot Output
![Screenshot 2024-11-17 074900](https://github.com/user-attachments/assets/f33dda59-f881-4bae-aa60-68c9f3375d7f)

#### Deskripsi Program
Program ini menentukan apakah suatu titik berada di dalam dua lingkaran yang diberikan oleh pengguna, berdasarkan koordinat pusat dan radius kedua lingkaran tersebut.

#### Algoritma Program
1. Masukkan koordinat pusat dan radius dari dua lingkaran.
2. Masukkan koordinat titik yang akan diperiksa.
3. Hitung jarak antara titik dan pusat lingkaran.
4. Periksa apakah jarak titik lebih kecil atau sama dengan radius lingkaran.
5. Outputkan hasil apakah titik berada di dalam salah satu, kedua, atau tidak ada lingkaran.

#### Cara Kerja
1. Fungsi jarak menghitung jarak antara titik pusat lingkaran dan titik yang diperiksa.
2. Fungsi didalam memeriksa apakah jarak titik ke pusat lingkaran lebih kecil atau sama dengan radius lingkaran.
3. Program kemudian mengevaluasi apakah titik berada di dalam salah satu atau kedua lingkaran, dan menampilkan hasilnya.

## 2. Program Pengolahan dan Analisis Data Array
#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// Menampilkan seluruh isi array
	fmt.Println("\nSeluruh isi array:", array)

	// Menampilkan elemen dengan indeks ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// Menampilkan elemen dengan indeks genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	// Menampilkan elemen dengan indeks kelipatan x
	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(array); i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	// Menghapus elemen pada indeks tertentu
	var index int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(array) {
		array = append(array[:index], array[index+1:]...)
		fmt.Println("Array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// Menampilkan rata-rata elemen array
	sum := 0
	for _, v := range array {
		sum += v
	}
	average := float64(sum) / float64(len(array))
	fmt.Printf("Rata-rata elemen array: %.2f\n", average)

	// Menampilkan standar deviasi elemen array
	var variance float64
	for _, v := range array {
		variance += math.Pow(float64(v)-average, 2)
	}
	variance /= float64(len(array))
	stdDev := math.Sqrt(variance)
	fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDev)

	// Menampilkan frekuensi suatu bilangan
	var target int
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scan(&target)
	count := 0
	for _, v := range array {
		if v == target {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, count)
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-17 075521](https://github.com/user-attachments/assets/2ee302f1-53fe-4608-b2f9-500259874325)

#### Screenshoot Output
![Screenshot 2024-11-17 075806](https://github.com/user-attachments/assets/53edf973-ac22-400a-a296-049820f0d53f)

#### Deskripsi Program
Program ini dibuat menggunakan bahasa Go untuk memanipulasi dan menganalisis data berupa array integer. Program menerima input jumlah elemen array dan nilai-nilai elemen array, kemudian memberikan berbagai informasi sesuai kebutuhan. Program ini mencakup operasi seperti menampilkan elemen array tertentu, menghitung rata-rata, standar deviasi, serta frekuensi elemen tertentu.

#### Algoritma Program 
1. Meminta pengguna memasukkan jumlah elemen array.
2. Membaca nilai elemen array satu per satu.
3. Menampilkan seluruh isi array.
4. Menampilkan elemen array berdasarkan:
   - Indeks ganjil.
   - Indeks genap.
   - Indeks kelipatan angka tertentu yang dimasukkan pengguna.
5. Menghapus elemen pada indeks tertentu sesuai input pengguna.
6. Menghitung rata-rata elemen array.
7. Menghitung standar deviasi elemen array.
8. Menghitung dan menampilkan frekuensi kemunculan bilangan tertentu dalam array.

#### Cara Kerja
1. Inisialisasi: Program dimulai dengan meminta jumlah elemen dan nilai-nilai array dari pengguna.
2. Operasi pada Array:
   - Menampilkan elemen sesuai kriteria tertentu (indeks ganjil, genap, atau kelipatan angka).
   - Menghapus elemen di indeks tertentu dan memperbarui array.
3. Perhitungan Statistik:
   - Rata-rata dihitung dengan membagi total elemen dengan jumlah elemen.
   - Standar deviasi dihitung menggunakan formula akar dari varians.
4. Frekuensi Elemen: Program menghitung berapa kali bilangan tertentu muncul dalam array.
5. Output: Hasil setiap operasi ditampilkan kepada pengguna secara bertahap.

## 3. Program Pencatat Pemenang Pertandingan Sepak Bola
```go
package main

import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var winners []string

	// Meminta input nama klub
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scan(&clubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scan(&clubB)

	fmt.Println("\nMasukkan skor untuk pertandingan (input skor negatif untuk berhenti):")

	// Loop untuk mencatat skor pertandingan
	match := 1
	for {
		fmt.Printf("Pertandingan %d - Skor %s: ", match, clubA)
		fmt.Scan(&scoreA)
		fmt.Printf("Pertandingan %d - Skor %s: ", match, clubB)
		fmt.Scan(&scoreB)

		// Hentikan jika skor tidak valid
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai.\n")
			break
		}

		// Menentukan hasil pertandingan
		if scoreA > scoreB {
			winners = append(winners, clubA)
			fmt.Printf("Hasil %d: %s menang\n", match, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
			fmt.Printf("Hasil %d: %s menang\n", match, clubB)
		} else {
			winners = append(winners, "Draw")
			fmt.Printf("Hasil %d: Draw\n", match)
		}
		match++
	}

	// Menampilkan daftar pemenang
	fmt.Println("Daftar hasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
	}
}

```
#### Screenshoot Source Code
![Screenshot 2024-11-17 080753](https://github.com/user-attachments/assets/b3ed3fbb-298d-4f04-b8c4-25cbc05c3979)

#### Screenshoot Output
![Screenshot 2024-11-17 080808](https://github.com/user-attachments/assets/9005a6f4-c0ee-478a-a39e-4f6796afc518)

#### Deskripsi Program
Program ini digunakan untuk mencatat hasil pertandingan sepak bola antara dua klub. Setiap pertandingan akan meminta pengguna untuk memasukkan nama klub dan skor masing-masing tim. Program secara otomatis menyimpan nama klub pemenang ke dalam daftar, atau mencatat hasil pertandingan sebagai draw jika skor sama. Proses input akan berhenti jika salah satu skor yang dimasukkan tidak valid (misalnya negatif).

#### Algoritma Program
1. Meminta pengguna memasukkan nama klub A dan klub B.
2. Memulai loop untuk mencatat skor pertandingan:
 - Meminta input skor klub A dan klub B.
 - Jika skor tidak valid (negatif), hentikan proses input.
 - Bandingkan skor:
   - Jika skor klub A lebih besar, tambahkan klub A ke daftar pemenang.
   - Jika skor klub B lebih besar, tambahkan klub B ke daftar pemenang.
   - Jika skor sama, tambahkan hasil draw ke daftar.
3. Setelah loop selesai, tampilkan daftar hasil pertandingan.

#### Cara Kerja
1. Inisialisasi: Program meminta input nama kedua klub yang bertanding.
2. Cek Validasi Skor: Program menerima input skor untuk kedua klub. Jika ditemukan skor negatif, proses berhenti.
3. Menyimpan Hasil Pertandingan:
   - Menyimpan nama klub pemenang ke dalam daftar jika salah satu menang.
   - Menyimpan hasil draw jika skor kedua klub sama.
4. Menampilkan Hasil: Setelah selesai, program menampilkan daftar semua hasil pertandingan yang telah tercatat.

## 4. Program Membalik Urutan Karakter Array dan Mengecek Palindrom
#### Source Code
```go
package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array
func isiArray(t *tabel, n *int) {
	fmt.Println("Masukkan karakter (akhiri dengan .):")
	*n = 0
	for {
		var ch rune
		fmt.Scanf("%c", &ch)

		// Hentikan input jika bertemu dengan titik atau melebihi batas array
		if ch == '.' || *n >= NMAX {
			break
		}

		t[*n] = ch
		*n++
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalik array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array adalah palindrom (mengabaikan spasi)
func palindrom(t tabel, n int) bool {
	left, right := 0, n-1
	for left < right {
		// Abaikan spasi di kiri
		for left < right && t[left] == ' ' {
			left++
		}
		// Abaikan spasi di kanan
		for left < right && t[right] == ' ' {
			right--
		}
		// Periksa apakah karakter tidak sama
		if t[left] != t[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Isi array
	isiArray(&tab, &n)

	// Cetak array asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Balik array dan cetak
	balikanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// Cek palindrom
	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}


```
#### Screenshoot Source Code
![Screenshot 2024-11-17 082044](https://github.com/user-attachments/assets/e4047363-51fb-46e8-a66b-19fc5e3d73aa)

#### Screenshoot Output
![Screenshot 2024-11-17 082022](https://github.com/user-attachments/assets/e6e3d83b-cd68-40b4-819c-16a537683139)

#### Deskripsi Program
Program ini menggunakan array untuk menyimpan kumpulan karakter yang dimasukkan pengguna. Program akan membalik urutan isi array tersebut dan memeriksa apakah karakter-karakter tersebut membentuk sebuah palindrom. Palindrom adalah teks yang sama jika dibaca dari depan maupun belakang, seperti kata "KATAK" atau "APA".

#### Algoritma Program
1. Inisialisasi: Membuat array untuk menyimpan karakter hingga batas maksimum.
2. Input Data: Membaca karakter yang dimasukkan oleh pengguna hingga batas maksimum atau hingga karakter titik ditemukan.
3. Cetak Data: Menampilkan isi array sesuai urutan input.
4. Balik Urutan: Membalik isi array menggunakan fungsi khusus.
5. Cek Palindrom: Membandingkan karakter asli dengan karakter setelah dibalik untuk memeriksa apakah teks adalah palindrom.
6. Output Hasil: Menampilkan hasil array terbalik dan status apakah array tersebut adalah palindrom.

#### Cara Kerja
1. Program meminta pengguna memasukkan teks berupa karakter hingga bertemu dengan titik (tanda berhenti).
2. Program mencetak teks sesuai urutan input.
3. Teks dibalik menggunakan prosedur pembalik array, lalu dicetak ulang.
4. Program memeriksa apakah teks tersebut palindrom dan mencetak hasilnya.

### Kesimpulan
Tipe bentukan, array, dan map adalah fitur penting dalam pemrograman untuk mengelola data secara efisien. Tipe bentukan seperti alias mempermudah penulisan kode, sementara struct mengelompokkan berbagai tipe data menjadi satu kesatuan yang terorganisir. Array dapat digunakan dalam bentuk statis untuk ukuran tetap atau slice untuk ukuran dinamis, memberikan fleksibilitas lebih saat runtime. Map memungkinkan penyimpanan data berbasis pasangan kunci-nilai, membuat akses data lebih fleksibel. Dengan menggabungkan struct dan array, data kompleks seperti data mahasiswa dapat dikelola secara terstruktur, sehingga memudahkan pemrograman dan pengolahan data.

## Referensi 
[1] Kernighan, B. W., & Ritchie, D. M. (1988). The C Programming Language (2nd ed.). Prentice Hall.

[2] Deitel, P., & Deitel, H. (2015). C How to Program (8th ed.). Pearson.

[3] Tanenbaum, A. S., & Bos, H. (2014). Modern Operating Systems (4th ed.). Pearson.

[4] Robbins, K. A., & Robbins, S. (2003). Unix Systems Programming: Communication, Concurrency, and Threads. Prentice Hall.

[5] Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. Addison-Wesley.

[6] Reinders, J. (2007). Intel Threading Building Blocks. O'Reilly Media.

[7] Weiss, M. A. (2013). Data Structures and Algorithm Analysis in C (3rd ed.). Pearson.
