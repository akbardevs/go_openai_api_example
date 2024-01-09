# Avatara Test

Penjelasan singkat tentang aplikasi .

## Kebutuhan

Pastikan Docker terpasang di sistem . Jika belum, dapat mengunduh dan menginstal Docker dari [situs web resmi Docker](https://www.docker.com/products/docker-desktop).

## Langkah-langkah Menjalankan Aplikasi

Ikuti langkah-langkah berikut untuk menjalankan aplikasi menggunakan Docker:

### 1. Kloning Repositori

Duplikat/"Cloning" repositori Git ke sistem lokal:

```bash
git clone https://github.com/akbardevs/go_openai_api_example
cd go_openai_api_example
```

### 2. Pengaturan .env

Pada file `.env` di direktori root aplikasi dan Ganti `your_openai_api_key` dengan kunci API OpenAI .

```env
OPENAI_API_KEY=your_openai_api_key
```

### 3. Membangun Docker Image

Bangun Docker image dari aplikasi :

```bash
docker build -t go_openai_api_example .
```

### 4. Menjalankan Container

Jalankan container dari image yang telah dibuat:

```bash
docker run -dp 8080:8080 go_openai_api_example
```

Pastikan port yang digunakan sesuai dengan yang didefinisikan dalam aplikasi .

### 5. Mengakses Aplikasi

Setelah container berjalan, dapat mengakses aplikasi melalui browser atau alat pengujian API seperti Postman di `http://localhost:8080`.

## Dokumentasi API

<!-- (Tambahkan dokumentasi API  di sini jika tersedia) -->
