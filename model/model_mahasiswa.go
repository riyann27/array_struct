package model

import "c/node"

var DaftarMahasiswa [10]node.Mahasiswa
var Kouta int = 10
var Index int = 0

func CreateMahasiswa(emp node.Mahasiswa) bool {
	if Index < Kouta {
		DaftarMahasiswa[Index] = emp
		Index = Index + 1
		return true

	}
	return false
}

func ReadMahasiswa() []node.Mahasiswa {
	return DaftarMahasiswa[0:Index]
}

func UpdateMahasiswa(emp node.Mahasiswa, npm string) bool {
	for i := 0; i < Index; i++ {
		if DaftarMahasiswa[i].Npm == npm {
			DaftarMahasiswa[i] = emp
			return true
		}
	}
	return false
}

func DeleteMahasiswa(npm string) bool {
	for i := 0; i < Index; i++ {
		if DaftarMahasiswa[i].Npm == npm {
			DaftarMahasiswa[i] = DaftarMahasiswa[Index-1]
			Index = Index - 1
			return true
		}
	}
	return false
}
