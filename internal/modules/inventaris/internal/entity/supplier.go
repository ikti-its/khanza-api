package entity

type Supplier struct {
	Id       int    `db:"id"`
	Nama     string `db:"nama"`
	Alamat   string `db:"alamat"`
	Telepon  string `db:"telepon"`
	Kota     string `db:"kota"`
	Bank     string `db:"bank"`
	Rekening string `db:"rekening"`
}
