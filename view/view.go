package view

import (
    "bufio"
    "c/model"
    "c/node"
    "fmt"
    "os"
    "strings"
)

func Insert() {
    var mhs node.Mahasiswa
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan NPM: ")
    mhs.Npm, _ = reader.ReadString('\n')
    mhs.Npm = strings.TrimSpace(mhs.Npm)

    fmt.Print("Masukkan Nama: ")
    mhs.Nama, _ = reader.ReadString('\n')
    mhs.Nama = strings.TrimSpace(mhs.Nama)

    fmt.Print("Masukkan Jalan: ")
    mhs.Alamat.Jalan, _ = reader.ReadString('\n')
    mhs.Alamat.Jalan = strings.TrimSpace(mhs.Alamat.Jalan)

    fmt.Print("Masukkan Kota: ")
    mhs.Alamat.Kota, _ = reader.ReadString('\n')
    mhs.Alamat.Kota = strings.TrimSpace(mhs.Alamat.Kota)

    fmt.Print("Masukkan No Telp: ")
    mhs.NoTelp, _ = reader.ReadString('\n')
    mhs.NoTelp = strings.TrimSpace(mhs.NoTelp)

    fmt.Print("Masukkan Email: ")
    mhs.Email, _ = reader.ReadString('\n')
    mhs.Email = strings.TrimSpace(mhs.Email)

    if model.CreateMahasiswa(mhs) {
        fmt.Println("Data Mahasiswa berhasil ditambahkan.")
    } else {
        fmt.Println("Data Mahasiswa gagal ditambahkan, kuota penuh.")
    }
}

func Views() {
    data := model.ReadMahasiswa()
    if len(data) == 0 {
        fmt.Println("Tidak ada data mahasiswa.")
        return
    }

    fmt.Println("\n=== Daftar Mahasiswa ===")
    for _, mhs := range data {
		fmt.Println("-> NPM           : ",mhs.Npm)
		fmt.Println("-> NAMA          : ",mhs.Nama)
		fmt.Println("-> ALAMAT        : ",mhs.Alamat.Jalan,mhs.Alamat.Kota,)
		fmt.Println("-> NOMOR TELEPON : ",mhs.NoTelp)
		fmt.Println("-> EMAIL         : ",mhs.Email)
		fmt.Println("===========================================")
        // fmt.Printf("NPM: %s, Nama: %s, Alamat: %s, %s, No Telp: %s, Email: %s\n",
        //     mhs.Npm, mhs.Nama, mhs.Alamat.Jalan, mhs.Alamat.Kota, mhs.NoTelp, mhs.Email)
    }
}

func Update() {
    var mhs node.Mahasiswa
    var npm string
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan NPM yang akan diupdate: ")
    npm, _ = reader.ReadString('\n')
    npm = strings.TrimSpace(npm)

    fmt.Print("Masukkan NPM baru: ")
    mhs.Npm, _ = reader.ReadString('\n')
    mhs.Npm = strings.TrimSpace(mhs.Npm)

    fmt.Print("Masukkan Nama baru: ")
    mhs.Nama, _ = reader.ReadString('\n')
    mhs.Nama = strings.TrimSpace(mhs.Nama)

    fmt.Print("Masukkan Jalan baru: ")
    mhs.Alamat.Jalan, _ = reader.ReadString('\n')
    mhs.Alamat.Jalan = strings.TrimSpace(mhs.Alamat.Jalan)

    fmt.Print("Masukkan Kota baru: ")
    mhs.Alamat.Kota, _ = reader.ReadString('\n')
    mhs.Alamat.Kota = strings.TrimSpace(mhs.Alamat.Kota)

    fmt.Print("Masukkan No Telp baru: ")
    mhs.NoTelp, _ = reader.ReadString('\n')
    mhs.NoTelp = strings.TrimSpace(mhs.NoTelp)

    fmt.Print("Masukkan Email baru: ")
    mhs.Email, _ = reader.ReadString('\n')
    mhs.Email = strings.TrimSpace(mhs.Email)

    if model.UpdateMahasiswa(mhs, npm) {
        fmt.Println("Data Mahasiswa berhasil diupdate.")
    } else {
        fmt.Println("Data Mahasiswa tidak ditemukan.")
    }
}

func Delete() {
    var npm string
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan NPM yang akan dihapus: ")
    npm, _ = reader.ReadString('\n')
    npm = strings.TrimSpace(npm)

    if model.DeleteMahasiswa(npm) {
        fmt.Println("Data Mahasiswa berhasil dihapus.")
    } else {
        fmt.Println("Data Mahasiswa tidak ditemukan.")
    }
}