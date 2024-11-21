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
  Yesika Widiyani / 2311102195 <br>
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

# UNGUIDED
## Unguided - 1
### Study Case
1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.</br>
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.</br>
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.</br>

### Source Code
![carbon (8)](https://github.com/user-attachments/assets/c0033956-4ab1-4623-8c3e-eeba5c6ce124)

### Screenshot Output
![image](https://github.com/user-attachments/assets/404dee73-a297-421d-bf5c-1cbd99806c1a)

### Deskripsi Program
Program ini digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program menerima input berupa jumlah anak kelinci (N) dan berat masing-masing anak kelinci dalam satuan bilangan riil. Program akan menentukan berat terkecil dan terbesar di antara data yang dimasukkan.

**Algoritma**
1. **Input Data**:
   - Baca jumlah anak kelinci (**N**).
   - Baca berat masing-masing anak kelinci (**N bilangan riil**).
2. **Inisialisasi**:
   - Tetapkan nilai awal untuk berat terkecil (`min`) dan terbesar (`max`) menggunakan nilai dari berat kelinci pertama.
3. **Iterasi**:
   - Untuk setiap berat kelinci, bandingkan nilai dengan `min` dan `max`:
     - Jika berat lebih kecil dari `min`, perbarui `min`.
     - Jika berat lebih besar dari `max`, perbarui `max`.
4. **Output**:
   - Cetak berat terkecil dan terbesar.

**Cara Kerja**
1. Program meminta input dari pengguna untuk jumlah anak kelinci dan berat masing-masing kelinci.
2. Dengan menggunakan array, program menyimpan semua berat kelinci.
3. Program mencari berat terkecil dan terbesar dengan memeriksa setiap elemen array satu per satu.
4. Hasil berupa berat terkecil dan terbesar ditampilkan di layar.

## Unguided - 2
### Study Case
2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.</br>
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil menyatakan banyaknya ikan yang akan dijual.</br>
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.</br>

### Source Code
![carbon (9)](https://github.com/user-attachments/assets/2f274d9c-1dc3-4eac-b8f1-d2deec03b4f1)

### Screenshot Output
![image](https://github.com/user-attachments/assets/9f88da5a-1f68-4ca3-9271-0d6ea4021160)

### Deskripsi Program
Program ini digunakan untuk menghitung berat ikan yang dimasukkan ke dalam wadah dan rata-rata berat di setiap wadah. Input terdiri dari jumlah ikan (**x**) dan kapasitas wadah (**y**), serta berat masing-masing ikan. Program menghasilkan total berat ikan di setiap wadah dan rata-rata berat ikan per wadah.

**Algoritma**
1. **Input Data**:
   - Masukkan jumlah ikan (**x**) dan kapasitas wadah (**y**).
   - Jika **x ≤ 0**, **y ≤ 0**, atau **y > x**, hentikan program dengan pesan error.
   - Masukkan **x** berat ikan dalam bentuk array.

2. **Inisialisasi Wadah**:
   - Hitung jumlah wadah, `jumlahWadah = ceil(x / y)`.
   - Buat array `totalWadah` untuk menyimpan berat total ikan di setiap wadah, panjang array adalah `jumlahWadah`.

3. **Hitung Berat Total di Setiap Wadah**:
   - Iterasi melalui array berat ikan:
     - Tentukan indeks wadah berdasarkan posisi ikan: `wadahKe = i / y`.
     - Tambahkan berat ikan ke wadah tersebut.

4. **Hitung Rata-Rata Berat Wadah**:
   - Hitung total semua berat ikan.
   - Bagikan total berat dengan `jumlahWadah` untuk mendapatkan rata-rata berat.

5. **Output Data**:
   - Tampilkan total berat di setiap wadah.
   - Tampilkan rata-rata berat ikan per wadah.

**Cara Kerja**
1. Program meminta input jumlah ikan (**x**) dan kapasitas wadah (**y**).
2. Berat ikan dimasukkan dalam urutan, lalu dibagi ke dalam wadah dengan kapasitas tertentu.
3. Berat ikan dalam wadah dihitung berdasarkan urutan input.
4. Setelah semua berat dihitung, program menampilkan total berat per wadah dan rata-rata berat ikan.

## Unguided - 3
### Study Case
3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dan data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:
![image](https://github.com/user-attachments/assets/f366966c-05af-4f34-be97-d0cd13de9204)
![image](https://github.com/user-attachments/assets/cdd62571-c69c-4092-8986-4de784f58615)

### Source Code
![carbon (10)](https://github.com/user-attachments/assets/4b944031-75f3-4a51-8cf6-ae352804c1a4)

### Screenshot Output
![image](https://github.com/user-attachments/assets/6448b278-c336-485c-9d5f-96cd45565f6f)

### Deskripsi Program
Program ini dirancang untuk mencatat berat badan balita di Posyandu, menghitung berat minimum, maksimum, dan rata-rata dari data yang diberikan. Menggunakan array untuk menyimpan data berat, program memproses input dari pengguna, lalu memanfaatkan subprogram untuk menghitung nilai-nilai tersebut secara efisien. Hasil akhirnya ditampilkan dalam format yang mudah dipahami oleh petugas.</br>

**Algoritma**  
1. Masukkan jumlah balita (`n`) dan pastikan valid.  
2. Input berat tiap balita ke dalam array.  
3. Inisialisasi `bMin` dan `bMax` dengan elemen pertama array.  
4. Iterasi array untuk:  
   - Memperbarui `bMin` jika ditemukan berat lebih kecil.  
   - Memperbarui `bMax` jika ditemukan berat lebih besar.  
5. Hitung total berat untuk mencari rata-rata (`total / n`).  
6. Tampilkan berat minimum, maksimum, dan rata-rata.  

**Cara Kerja Singkat**  
1. Masukkan jumlah balita dan berat masing-masing ke array.  
2. Inisialisasi `bMin` dan `bMax` dari berat pertama.  
3. Iterasi array untuk memperbarui `bMin` dan `bMax`.  
4. Hitung total berat dan rata-rata.  
5. Tampilkan berat minimum, maksimum, dan rata-rata.
