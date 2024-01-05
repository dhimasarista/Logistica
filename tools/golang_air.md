Untuk menginstall `air` di Go (Golang), Anda perlu mengikuti langkah-langkah berikut:

1. Pastikan Go sudah terinstall di sistem Anda. Jika belum, Anda dapat mengunduh dan menginstall Go dari situs resmi: [https://golang.org/dl/](https://golang.org/dl/).

2. Buka terminal atau command prompt dan ketik perintah berikut untuk menginstall `air` menggunakan perintah `go install` bukan `go get`:

    ```bash
    go install github.com/cosmtrek/air@latest
    ```

    Perintah ini akan mengunduh dan menginstall `air` dari repositori GitHub.

3. Setelah proses instalasi selesai, Anda dapat menggunakan `air` untuk mengelola hot-reloading pada proyek Go Anda.

4. Sebelum menggunakan `air`, pastikan bahwa proyek Go Anda memiliki file konfigurasi `.air.toml` di direktori proyek. Jika tidak, Anda dapat membuatnya secara manual atau dengan menjalankan perintah:

    ```bash
    air init
    ```

    Perintah ini akan membuat file `.air.toml` yang dapat Anda sesuaikan sesuai kebutuhan proyek Anda.

5. Jalankan proyek Anda menggunakan `air`:

    ```bash
    air
    ```

    Ini akan memulai server pengembangan dan otomatis akan melakukan hot-reloading setiap kali ada perubahan pada file proyek Anda.

Pastikan untuk membaca dokumentasi `air` dan mengonfigurasi file `.air.toml` sesuai dengan kebutuhan proyek Anda. Dokumentasi `air` dapat ditemukan di [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air).