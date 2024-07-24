package entity

import (
	"github.com/google/uuid"
)

type GudangBarang struct {
	IdBarangMedis uuid.UUID `db:"id_barang_medis"`
	IdRuangan     int       `db:"id_ruangan"`
	Stok          int       `db:"stok"`
	NoBatch       string    `db:"no_batch"`
	NoFaktur      string    `db:"no_faktur"`
}
