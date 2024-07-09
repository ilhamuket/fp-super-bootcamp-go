# REST API "news" menggunakan Golang

![Logo Golang](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

## URL Live

Anda dapat mengakses proyek ini secara live di [https://news-golang-api-670f44c77a36.herokuapp.com/](https://news-golang-api-670f44c77a36.herokuapp.com/).

## Pengenalan

Assalamualaikum, nama saya Muhammad Ilham, peserta kelas Fullstack Sanber Super Bootcamp Jabar 2024 dengan fokus pada Golang dan Next.js. Saya senang memperkenalkan proyek REST API "news" yang telah saya buat menggunakan Golang. Proyek ini memungkinkan Anda untuk mengelola berita melalui API dengan fitur-fitur terstruktur.

## Fitur-fitur

- **Registrasi dan Otentikasi Pengguna**: Pengguna dapat mendaftar dan melakukan login untuk mengelola berita.
- **Manajemen Berita**: Pengguna yang terotentikasi dapat membuat, mengedit, dan menghapus berita.
- **Komentar**: Pengguna dapat menambahkan komentar pada berita yang ada.
- **Hak Akses Berbasis Peran**: Berbagai peran seperti admin dan editor memiliki hak akses yang berbeda pada operasi berita.

## Setup Proyek

Untuk menjalankan proyek ini, pastikan Anda memiliki lingkungan pengembangan Golang yang sudah siap. Clone repositori ini dan jalankan `go run main.go` untuk memulai server lokal.

## API Routes

### Registrasi dan Otentikasi

- **POST /register**
  - Deskripsi: Registrasi pengguna baru.
  
- **POST /login**
  - Deskripsi: Login pengguna untuk mendapatkan token JWT.

### Manajemen Berita

- **GET /news/:id**
  - Deskripsi: Mendapatkan detail berita berdasarkan ID.
  
- **GET /news**
  - Deskripsi: Mendapatkan semua berita yang tersedia.
  
- **POST /news**
  - Deskripsi: Membuat berita baru (memerlukan token JWT).
  
- **PUT /news/:id**
  - Deskripsi: Mengupdate berita yang ada (memerlukan token JWT dan peran admin/editor).
  
- **DELETE /news/:id**
  - Deskripsi: Menghapus berita (memerlukan token JWT dan peran admin).

### Manajemen Komentar

- **POST /comments**
  - Deskripsi: Menambahkan komentar pada berita.
  
- **GET /comments/:id**
  - Deskripsi: Mendapatkan detail komentar berdasarkan ID.
  
- **GET /news/comments/:news_id**
  - Deskripsi: Mendapatkan semua komentar berdasarkan ID berita.
  
- **PUT /comments/:id**
  - Deskripsi: Mengupdate komentar yang ada.
  
- **DELETE /comments/:id**
  - Deskripsi: Menghapus komentar.

### Endpoints lainnya

- **PUT /change-password**
  - Deskripsi: Mengganti password pengguna (memerlukan token JWT).
  
- **GET /profile**
  - Deskripsi: Mendapatkan profil pengguna (memerlukan token JWT).
  
- **PUT /profile**
  - Deskripsi: Mengupdate profil pengguna (memerlukan token JWT).
  
- **GET /users**
  - Deskripsi: Mendapatkan semua pengguna (memerlukan peran admin).
  
- **GET /users/:id**
  - Deskripsi: Mendapatkan detail pengguna berdasarkan ID (memerlukan peran admin).

## Dokumentasi API

Dokumentasi API menggunakan Swagger dan dapat diakses melalui endpoint `/swagger`. Ini memberikan informasi lebih lanjut tentang setiap endpoint dan parameter yang dibutuhkan.

## Catatan

Pastikan untuk mengonfigurasi database dan lingkungan proyek sesuai dengan kebutuhan Anda sebelum menjalankannya di lingkungan produksi.

## Kontribusi

Anda dapat berkontribusi pada proyek ini dengan mengirimkan pull request atau melaporkan masalah yang ditemukan.

Terima kasih telah membaca tentang proyek REST API "news" saya yang dibangun dengan Golang. Semoga bermanfaat!
