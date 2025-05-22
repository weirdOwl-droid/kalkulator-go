Berikut adalah catatan penjelasan untuk kode program **kalkulator-go** yang dapat dimasukkan ke dalam README proyek GitHub:

---

# 🧮 **Kalkulator CLI dengan Go**
Ini adalah program **Kalkulator berbasis CLI** yang ditulis dalam bahasa **Go** untuk melakukan operasi matematika sederhana seperti penjumlahan, pengurangan, perkalian, dan pembagian.

## 🚀 **Fitur Utama**
✅ Penjumlahan (`+`)  
➖ Pengurangan (`-`)  
✖️ Perkalian (`*`)  
➗ Pembagian (`/`) dengan validasi pembagian nol  

## 🏗 **Struktur Kode**
Program ini menggunakan **Go** dengan package bawaan:
- `fmt`: Menampilkan teks ke terminal dan menerima input pengguna.

## 📜 **Penjelasan Kode**
- **`fmt.Scanln(&num1)`** → Mengambil input angka pertama dari pengguna.
- **`fmt.Scanln(&operator)`** → Mengambil input operator matematika (`+`, `-`, `*`, `/`).
- **`fmt.Scanln(&num2)`** → Mengambil input angka kedua dari pengguna.
- **`switch-case`** → Mengevaluasi operator dan menghitung hasil sesuai operasi yang dipilih.
- **Validasi pembagian oleh nol** → Mencegah error dengan kondisi `if num2 == 0`.

## 🔧 **Cara Menjalankan**
Pastikan Go sudah terinstall, lalu jalankan:
```bash
go run main.go
```
Atau build terlebih dahulu dengan:
```bash
go build -o kalkulator main.go
./kalkulator
```

---
