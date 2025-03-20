package main

import (
    "c/view" // Pastikan path import ini sesuai dengan struktur folder Anda
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings" // Import package strings untuk menggunakan strings.TrimSpace
)

func menuUtama() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\n=== MENU UTAMA ===")
        fmt.Println("1. Tambah Data Mahasiswa")
        fmt.Println("2. Tampilkan Data Mahasiswa")
        fmt.Println("3. Update Data Mahasiswa")
        fmt.Println("4. Hapus Data Mahasiswa")
        fmt.Println("5. Keluar")
        fmt.Print("Pilih menu: ")

        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input) // Menghapus newline dan spasi

        choice, err := strconv.Atoi(input) // Konversi input ke integer
        if err != nil {
            fmt.Println("Input tidak valid, silakan masukkan angka.")
            continue
        }

        switch choice {
        case 1:
            fmt.Println("Anda memilih: Tambah Data Mahasiswa")
            view.Insert() // Panggil fungsi Insert dari package view
        case 2:
            fmt.Println("Anda memilih: Tampilkan Data Mahasiswa")
            view.Views() // Panggil fungsi Views dari package view
        case 3:
            fmt.Println("Anda memilih: Update Data Mahasiswa")
            view.Update() // Panggil fungsi Update dari package view
        case 4:
            fmt.Println("Anda memilih: Hapus Data Mahasiswa")
            view.Delete() // Panggil fungsi Delete dari package view
        case 5:
            fmt.Println("Terima kasih! Program selesai.")
            return // Keluar dari program
        default:
            fmt.Println("Pilihan tidak valid, silakan coba lagi.")
        }
    }
}

func main() {
    menuUtama() // Jalankan menu utama
}