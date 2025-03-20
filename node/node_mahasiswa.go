package node

type Adress struct {
	Jalan, Kota string
}

type Mahasiswa struct {
	Npm, Nama string
	Alamat    Adress
	NoTelp    string
	Email     string
}
