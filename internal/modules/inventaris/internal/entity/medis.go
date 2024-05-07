package entity

import "github.com/google/uuid"

type Medis struct {
	Id      uuid.UUID `db:"id"`
	Nama    string    `db:"nama"`
	Jenis   string    `db:"jenis"`
	Satuan  int       `db:"id_satuan"`
	Harga   float64   `db:"harga"`
	Stok    int       `db:"stok"`
	Updater uuid.UUID `db:"updater"`
}
