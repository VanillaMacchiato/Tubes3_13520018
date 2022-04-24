# Backend tubes 3

Menjalankan backend pada local:

1. Install Go
2. Buka CLI (dan arahkan ke directory ini), jalankan `go get` pada CLI
3. Buat file `.env` yang di-copy dari `.env.sample`, isi DATABASE_URL dan lain-lain
4. Jalankan `go run main.go` pada CLI

# API Endpoint

- [Menambahkan Penyakit pada Database](#menambahkan-penyakit-pada-database)
- [Melakukan Prediksi pada Pasien](#melakukan-prediksi-pada-pasien)

## Menambahkan Penyakit pada Database

**Request**

`POST /api/v1/add-disease`

Request dengan multipart/form-data berisi dua pair key-value sebagai berikut

| Key          | Value                    |
| ------------ | ------------------------ |
| file         | [isi dengan file]        |
| disease-name | isi dengan nama penyakit |

Contoh:

| Key          | Value   |
| ------------ | ------- |
| file         | dna.txt |
| disease-name | LIGMA   |

File yang diunggah melalui multipart form-data pada key `file` ini berisi raw text DNA. Jika mengandung karakter yang tidak diinginkan (bukan ACGT), penambahan ke database ditolak.

## Melakukan Prediksi pada Pasien

**Request**

`POST /api/v1/predict-patience`

Request dengan multipart/form-data berisi tiga pair key-value sebagai berikut

| Key       | Value                    |
| --------- | ------------------------ |
| file      | [isi dengan file DNA]    |
| name      | nama pasien              |
| disease   | prediksi nama penyakit   |
| algorithm | algoritma yang digunakan |

File yang diunggah melalui multipart form-data pada key `file` ini berisi raw text DNA. Jika mengandung karakter yang tidak diinginkan (bukan ACGT), penambahan ke database ditolak.

## Mendapatkan Hasil Prediksi

**Request**

`GET /api/v1/result`

| Parameter | Format        |
| --------- | ------------- |
| date      | YYYY-MM-DD    |
| name      | Nama penyakit |

Mengembalikan hasil pasien yang telah diprediksi.

Contoh:

- `/api/v1/result?date=2022-04-22`
- `/api/v1/result?name=HIV`
- `/api/v1/result?name=HIV&date=2022-04-22`
