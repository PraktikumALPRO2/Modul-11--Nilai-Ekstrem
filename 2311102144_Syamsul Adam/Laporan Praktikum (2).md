# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 11</h2>
# <h2 align="center">PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Syamsul Adam - 2311102144</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI

### 1. Ide Pencarian Nilai Max/Min
Pencarian adalah suatu proses yang lazim dilakukan di dalam kehidupan sehari-hari. Contoh penggunaannya dalam kehidupan nyata sangat beragam, misalnya pencarian file di dalam directory komputer, pencarian suatu teks di dalam sebuah dokumen, pencarian buku pada rak buku, dan contoh lainnya. Pertama pada modul ini akan dipelajari salah satu algoritma pencarian nilai terkecil atau terbesar pada sekumpulan data, atau biasa disebut pencarian nilai ekstrim.<br/>

Ide algoritma sederhana sekali. Karena data harus diproses secara sekuensial, maka nilai atau indeks ke nilai maksimum dari data yang telah diproses disimpan untuk dibandingkan dengan data berikutnya. Nilai yang berhasil disimpan sampai algoritma tersebut berakhir adalah nilai maksimum yang dicari. Adapun algoritmanya secara umum adalah sebagai berikut:<br/>
    1. Jadikan data pertama sebagai nilai ekstrim<br/>
    2. Lakukan validasi nilai ekstrim dari data kedua hingga data terakhir.<br/>
     - Apabila nilai ekstrim tidak valid, maka update nilai ekstrim tersebut dengan data
     yang dicek.<br/>
    3. Apabila semua data telah dicek, maka nilai ekstrim yang dimiliki adalah valid.<br/>
Berikut ini adalah notasi dalam pseudocode dan bahasa Go, misalnya untuk pencarian nilai terbesar atau maksimum:<br/>
![image](https://github.com/user-attachments/assets/e5c4bd01-f629-4951-b79e-2860152b9a01)<br/>

### 2. Pencarian Nilai Ekstrim pada Array Bertipe Data Besar
Misalnya terdefinisi sebuah array of integer dengan kapasitas 2023, dan array terisi sejumlah N bilangan bulat, kemudian pencarian nilai terkecil dilakukan pada array tersebut. Perhatikan potongan program dalam bahasa Go berikut ini:<br/>
![image](https://github.com/user-attachments/assets/807e5b5b-844f-48cd-8824-dbfb03746334)<br/>
Potongan program di atas sedikit berbeda dengan sebelumnya karena penggunaan indeks array pada bahasa Go di mulai dari nol atau "0" seperti penjelasan pada modul 7. Selanjutnya, pada penjelasan di awal bab 3 telah disampaikan bahwa pada pencarian yang terpenting adalah posisi atau indeks dari nilai yang dicari dalam kumpulan data atau array. Oleh karena itu modifikasi pada program di atas dapat dilihat pada potongan program berikut ini!<br/>
![image](https://github.com/user-attachments/assets/5c6d92a8-0d0b-4750-b9af-2caf42ce084e)<br/>

### 3. Pencarian Nilai Ekstrim pada Array Bertipe Data Terstruktur
Pada kasus yang lebih kompleks pencarian ekstrim dapat juga dilakukan, misalnya mencari data mahasiswa dengan nilai terbesar, mencari lagu dengan durasi terlama, mwncari pembalap yang memiliki catatan waktu balap tercepat, dan sebagainya. Sebagai contoh misalnya terdapat array yang digunakan untuk menyimpan data mahasiswa, kemudian terdapat fungsi IPK yang digunakan untuk encari data mahasiswa dengan IPK tertinggi.<br/>
![image](https://github.com/user-attachments/assets/697c1778-2074-424e-b844-1efba0a99f27)<br/>
Apabila diperhatikan potongan program diatas, maka kita akan memperoleh ipk tertinggi, tetapi kita tidak memperoleh identitas mahasiswa dengan ipk tertinggi tersebut. Maka seperti penjelasan yang sudah diberikan sebelumnya, maka pencarian yang dilakukan bisa mengembalikan indeks mahasiswa dengan ipk tertinggi tersebut. Berikut ini adalah modifikasinya!<br/>
![image](https://github.com/user-attachments/assets/0612778a-1ed6-462e-ac58-7dde4c288bd9)<br/>
Sehingga melalui algoritma di atas, identitas mahasiswa dapat diperoleh, misalnya T[idx].nama, T[idx].nim, T[idx].kelas, hingga T[idx].jurusan.<br/>

## II. UNGUIDED

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. <br/> Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.<br/> Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

```go
package main

import "fmt"

func main() {
	// Input
	var n int
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scanln(&n)
	weights := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&weights[i])
	}

	// Mencari berat kelinci terkecil dan terbesar
	minWeight := weights[0]
	maxWeight := weights[0]
	for i := 1; i < n; i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	// Output
	fmt.Println("Berat kelinci terkecil:", minWeight)
	fmt.Println("Berat kelinci terbesar:", maxWeight)
}
```
## Output:![Screenshot 2024-11-24 224245](https://github.com/user-attachments/assets/fccc20d3-d5c9-4b7c-98fb-9bc1f1404548)


Program di atas untuk menghitung dan menampilkan berat terkecil dan terbesar dari sekumpulan data berat kelinci.  Setelah jumlah kelinci dan berat masing-masing dimasukkan, program membandingkan tiap berat untuk menentukan nilai minimum dan maksimum, lalu menampilkan hasilnya 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. <br/> Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukkan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.<br/> Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

```go
package main

import (
	"fmt"
)

func main() {
	// Input

	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	var beratIkan []float64
	fmt.Println("Masukkan berat ikan (x bilangan riil):")
	for i := 0; i < x; i++ {
		var berat float64
		fmt.Scanln(&berat)
		beratIkan = append(beratIkan, berat)
	}

	// Hitung total berat ikan di setiap wadah
	totalBerat := make([]float64, y)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
		

		totalBerat[i%y] += beratIkan[i]
	}

	// Hitung berat rata-rata ikan di setiap wadah
	rataRataBerat := 0.0
	for i := 0; i < y; i++ {
		rataRataBerat += totalBerat[i]
	}
	rataRataBerat /= float64(y)

	// Output: Dua baris, baris pertama berisi total berat, baris kedua berisi rata-rata
	fmt.Println("Total berat ikan di setiap wadah:")
	for i := 0; i < y; i++ {
		fmt.Printf("%.2f ", totalBerat[i])
	}
	fmt.Println()
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", rataRataBerat)
}

```
## Output:![Screenshot 2024-11-24 213248](https://github.com/user-attachments/assets/e89fed4c-5656-4ef0-9ac3-ba71e8a54010)

Program ini berguna untuk menghitung berat ikan ke dalam beberapa wadah berdasarkan kapasitas yang ditentukan. Program ini dapat digunakan dalam konteks perikananterkait distribusi berat ikan dalam wadah. Program ini berguna tentang total dan rata-rata berat ikan di setiap wadah.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan rataratanya.<br/> Buatlah program dengan spesifikasi subprogram sebagai berikut:<br/>
![image](https://github.com/user-attachments/assets/5519ae3c-963f-4e37-a129-e1d32f90879f)<br/>

go
Copy code


```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
    *bMin = arrBerat[0]
    *bMax = arrBerat[0]
    for i := 1; i < n; i++ { // Hanya iterasi hingga n
        if arrBerat[i] < *bMin {
            *bMin = arrBerat[i]
        }
        if arrBerat[i] > *bMax {
            *bMax = arrBerat[i]
        }
    }
}

func rerata(arrBerat arrBalita, n int) float64 {
    var sum float64
    for i := 0; i < n; i++ { // Hanya iterasi hingga n
        sum += arrBerat[i]
    }
    return sum / float64(n)
}

func main() {
    var n int
    fmt.Print("Masukan banyak data berat balita: ")
    fmt.Scanln(&n)
    var arrBalita arrBalita
    for i := 0; i < n; i++ {
        fmt.Printf("Masukan berat balita ke-%d: ", i+1)
        fmt.Scanln(&arrBalita[i])
    }

    var bMin, bMax float64
    hitungMinMax(arrBalita, n, &bMin, &bMax) // Pass n ke fungsi
    fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
    fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
    fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(arrBalita, n)) // Pass n ke fungsi
}

```
## Output:![Screenshot 2024-11-24 223738](https://github.com/user-attachments/assets/502093ee-774c-42e8-8c21-715f1026c899)


Program ini efektif dalam menghitung dan menampilkan berat minimum, maksimum, dan rata-rata dari data berat balita. Program ini dapat digunakan dalam berbagai konteks. Dengan menggunakan fungsi terpisah untuk perhitungan, program ini juga menunjukkan struktur yang baik dan memudahkan pemeliharaan serta pengembangan lebih lanjut.

## KESIMPULAN
Secara keseluruhan, ketiga program ini menunjukkan kemampuan untuk mengolah data dengan cara yang terstruktur. Masing-masing program memiliki fokus yang berbeda, tetapi semuanya memberikan hasil yang bermanfaat berdasarkan input pengguna. Penggunaan array dan fungsi dalam program-program ini juga mencerminkan praktik pemrograman yang baik, yang dapat memudahkan pengembangan dan pemeliharaan. Program-program ini dapat diterapkan dalam berbagai konteks, termasuk peternakan, perikanan, dan kesehatan,
## Daftar Pustaka
Modul 11 Praktikum algoritma pemograman.
