# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 8</h2>
# <h2 align="center">PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
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

## 2. UNGUIDED

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. <br/> Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.<br/> Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

```go
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 {
		fmt.Println("Jumlah kelinci harus lebih dari 0.")
		return
	}

	weights := make([]float64, n)
	fmt.Printf("Masukkan berat masing-masing kelinci (%d angka):\n", n)

	// Input data berat kelinci
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	// Inisialisasi nilai minimum dan maksimum
	minWeight := weights[0]
	maxWeight := weights[0]

	// Cari berat terkecil dan terbesar
	for _, weight := range weights {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	// Tampilkan hasil
	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/1a6b7b7f-4fde-4076-bfad-c1665172af3d)

Program di atas mencari berat kelinci terkecil dan terbesar dari data yang diinput. Setelah jumlah kelinci dan berat masing-masing dimasukkan, program membandingkan tiap berat untuk menentukan nilai minimum dan maksimum, lalu menampilkan hasilnya dalam format dua desimal.

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. <br/> Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukkan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.<br/> Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukkan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi input
	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
		return
	}

	weights := make([]float64, x)
	fmt.Printf("Masukkan berat %d ikan:\n", x)

	// Input berat ikan
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	// Proses pembagian ikan ke dalam wadah
	totalWadah := int(math.Ceil(float64(x) / float64(y)))
	wadah := make([]float64, totalWadah)
	avgWeight := 0.0

	for i := 0; i < x; i++ {
		wadah[i/y] += weights[i]
		avgWeight += weights[i]
	}

	// Hitung rata-rata
	avgWeight /= float64(totalWadah)

	// Output hasil
	fmt.Println("Berat total di setiap wadah:")
	for i := 0; i < totalWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()

	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", avgWeight)
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/30572733-e9cb-4185-a631-f07a9a3e5e16)

Program di atas membagi ikan ke dalam wadah sesuai kapasitas, menghitung total berat ikan di setiap wadah, dan menentukan rata-rata berat per wadah. Data jumlah ikan, kapasitas wadah, dan berat masing-masing ikan diinputkan, lalu program memprosesnya berdasarkan urutan dan menampilkan hasilnya.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan rataratanya.<br/> Buatlah program dengan spesifikasi subprogram sebagai berikut:<br/>
![image](https://github.com/user-attachments/assets/5519ae3c-963f-4e37-a129-e1d32f90879f)<br/>

go
Copy code


```go
package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin, *bMax = arrBerat[0], arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func hitungRerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rerata := hitungRerata(arrBerat, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
```
## Output:<br/> ![image](https://github.com/user-attachments/assets/d2cdb4d4-7ac0-44f9-8a8f-0a599a7b4e51)

Program di atas mencatat berat balita, lalu menghitung nilai minimum, maksimum, dan rata-rata menggunakan dua fungsi terpisah, kemudian menampilkan hasilnya.



