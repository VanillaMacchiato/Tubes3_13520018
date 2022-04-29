<h1 align="center">
  <br>
  <a><img src="https://media.discordapp.net/attachments/893484082275708980/969508968487600189/unknown.png?width=543&height=701" width="200"></a>
  <br>
 DNA At Work
  <br>
</h1>

<h4 align="center">A simple website for DNA pattern matching</h4>
<p align="center">
  <a href=https://dna-at-work.pages.dev>
    <img src="https://media.discordapp.net/attachments/893484082275708980/969515386141802506/badge.png"
         alt="Website Link">
  </a>
</p>
<p align="center">
  <a href="#description">Description</a> •
  <a href="#features">Features</a> •
  <a href="#hosting-locally">Hosting Locally</a> •
  <a href="#credits">Credits</a> •
  <a href="#footnote">Footnote</a> 
</p>

## Description
Sebuah aplikasi web simpel DNA Pattern Matching.  Dengan memanfaatkan algoritma  String Matching (Boyer-Moore serta Knuth-Morris-Pratt) dan Regular Expression telah dibagin sebuah aplikasi  interaktif untuk mendeteksi apakah seorang pasien mempunyai penyakit genetik tertentu. Hasil  prediksi tersebut  disimpan  pada  basis  data  untuk  kemudian  dapat  ditampilkan  berdasarkan query pencarian.

## Features
### Halaman Utama
![Gambar Halaman Utama](https://media.discordapp.net/attachments/893484082275708980/969510210358116372/unknown.png?width=1247&height=702)
### Input Penyakit
![Gambar Halaman Input Penyakit](https://media.discordapp.net/attachments/893484082275708980/969510210043527198/unknown.png?width=1247&height=702)
### Prediksi Penyakit
![Gambar Halaman Prediksi Penyakit](https://media.discordapp.net/attachments/893484082275708980/969510209687023656/unknown.png?width=1247&height=702)
### Hasil Prediksi
![Gambar Halaman prediksi Penyakit](https://media.discordapp.net/attachments/893484082275708980/969510209380814868/unknown.png?width=1247&height=702)

## Hosting Locally
Pada sisi *Backend*, aplikasi web ini membutuhkan [Go](https://go.dev/dl/) untuk dijalankan. Dengan asumsi Anda sudah mengunduh semua source code pada repository ini, cara menjalankan *backend* diberikan sebagai berikut:
1. Install [Go](https://go.dev/dl/)
2. Buka folder yang mengandung source code repository.
3. Buka folder src, kemudian folder backend
4. Ubah nama file `.env.sample` ke `.env` dan ubah isinya juga menggunakan text editor
```
DATABASE_URL={LINK_MONGODB_DATABASE}
FE_URL=*
```
5. Buka terminal pada folder tersebut
6. Install Dependencies 
```bash
# Install Dependencies
go get
```
7. Jalankan main.go
```bash
# Initiate backend
go run main.go
```
8. Apabila Anda di Windows, Anda perlu menyetujui *popup* firewall 
9. Selesai!

Pada sisi *Frontend*, aplikasi web ini membutuhkan VueCLI dan NodeJS untuk dijalankan. Dengan asumsi Anda sudah mengunduh semua source code pada repository ini, cara menjalankan *frontend* diberikan sebagai berikut:
1. Install [npm](https://docs.npmjs.com/cli/v7/configuring-npm/install)
2. Buka folder yang mengandung source code repository.
3. Buka folder src, kemudian folder frontend
4. Buka terminal pada folder tersebut
5. Install NodeJS dan VueCLI
```bash
# Install NodeJS & VueCLI
npm install nodejs @vue/cli
```
6.  Install Dependencies lainnya
``` bash
# Install Dependencies
npm install
```
7. Jalankan program
``` bash
# Initiate Frontend
npm run serve
```
8. Buka alamat yang terdapat pada CLI (biasanya [localhost:8081](localhost:8081))
9. Selesai!

## Credits
<table>
    <tr>
      <td><b>Nama</b></td>
      <td><b>NIM</b></td>
    </tr>
    <tr>
      <td><b>Bariza Haqi</b></td>
      <td><b>13520018</b></td>
    </tr>
        <tr>
      <td><b>Vito Ghifari</b></td>
      <td><b>13520153</b></td>
    </tr>
        <tr>
      <td><b>Farrel Farandieka Fibriyanto</b></td>
      <td><b>13520054</b></td>
    </tr>
</table>

## Footnote
- Setup yang disediakan pada [Hosting Locally](#hosting-locally) berasumsikan belum ada port yang sedang berjalan pada mesin Anda dan bahwa backend dijalankan terlebih dahulu, baru kemudian frontend. Apabila kondisi di atas tidak terpenuhi, silahkan ganti port di *frontend* pada source code
- Dibuat atas permintaan [Tugas Besar 3 IF2211](https://informatika.stei.itb.ac.id/~rinaldi.munir/Stmik/2021-2022/Tugas-Besar-3-IF2211-Strategi-Algoritma-2022.pdf) Teknik Informatika ITB tahun ajaran 2021/2022
