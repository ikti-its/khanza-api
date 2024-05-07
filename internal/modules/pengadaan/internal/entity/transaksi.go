package entity

import "github.com/google/uuid"

type Transaksi struct {
	Id      uuid.UUID `db:"id"`
	IdMedis uuid.UUID `db:"id_barang_medis"`
	Batch   string    `db:"no_batch"`
	Faktur  string    `db:"no_faktur"`
	Jumlah  int       `db:"jumlah_keluar"`
	Updater uuid.UUID `db:"updater"`
}
