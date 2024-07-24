package entity

type Role struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type Jabatan struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type Departemen struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type StatusAktif struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}

type Shift struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}

type AlasanCuti struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}

type IndustriFarmasi struct {
	Id      int    `db:"id"`
	Kode    string `db:"kode"`
	Nama    string `db:"nama"`
	Alamat  string `db:"alamat"`
	Kota    string `db:"kota"`
	Telepon string `db:"telepon"`
}

type SatuanBarangMedis struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type JenisBarangMedis struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type KategoriBarangMedis struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type GolonganBarangMedis struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type Ruangan struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type SupplierBarangMedis struct {
	Id         int    `db:"id"`
	Nama       string `db:"nama"`
	Alamat     string `db:"alamat"`
	NoTelp     string `db:"no_telp"`
	Kota       string `db:"kota"`
	NamaBank   string `db:"nama_bank"`
	NoRekening string `db:"no_rekening"`
}
