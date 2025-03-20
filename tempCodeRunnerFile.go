package main

import (
	"fmt"
	"c/model"
	"c/node"
)

func main() {
	mahasiswa1 := node.Mahasiswa{
		Npm:   "06.2024.1.07777",
		Nama:  "Riyan",
		Alamat: node.Adress{"Jl.Kedung Baruk", " Rungkut "},
		NoTelp: "085159400775",
		Email:  "aku@com",
	}
	mahasiswa2 := node.Mahasiswa{
		Npm:   "07777",
		Nama:  "yan",
		Alamat: node.Adress{"Jl.Kedung Baruk", " Rungkut "},
		NoTelp: "775",
		Email:  "saya@com",
	}
	model.CreateMahasiswa(mahasiswa1)
	model.CreateMahasiswa(mahasiswa2)

	for _, emp := range model.ReadMahasiswa() {
		fmt.Println("-> NPM     : ", emp.Npm)
		fmt.Println("-> NAMA    : ", emp.Nama)
		fmt.Println("-> Alamat  : ", emp.Alamat)
		fmt.Println("-> Telepon : ", emp.NoTelp)
		fmt.Println("-> EMAIL   : ", emp.Email)
	}

	// Update testing
	updatedMahasiswa := node.Mahasiswa{
		Npm:   "07777",
		Nama:  "any",
		Alamat: node.Adress{"Jl.Kedung Baruk", " Rungkut "},
		NoTelp: "775",
		Email:  "saya@com",
	}
	model.UpdateMahasiswa(updatedMahasiswa, 1) // Angka 1 adalah indeks mahasiswa yang ingin diupdate

	fmt.Println("After Update")
	for _, emp := range model.ReadMahasiswa() {
		fmt.Println("-> NPM     : ", emp.Npm)
		fmt.Println("-> NAMA    : ", emp.Nama)
		fmt.Println("-> Alamat  : ", emp.Alamat)
		fmt.Println("-> Telepon : ", emp.NoTelp)
		fmt.Println("-> EMAIL   : ", emp.Email)
	}
}