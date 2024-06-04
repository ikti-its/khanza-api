package entity

import "github.com/google/uuid"

type Medis struct {
	Id          uuid.UUID `db:"id"`
	Nama        string    `db:"nama"`
	Jenis       string    `db:"jenis"`
	Satuan      int       `db:"id_satuan"`
	Harga       float64   `db:"harga"`
	Stok        int       `db:"stok"`
	StokMinimum int       `db:"stok_minimum"`
	Notifikasi  int       `db:"notifikasi_kadaluwarsa_hari"`
	Updater     uuid.UUID `db:"updater"`
}
