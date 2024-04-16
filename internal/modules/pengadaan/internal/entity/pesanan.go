package entity

import (
	"github.com/google/uuid"
	"time"
)

type Pesanan struct {
	Id          uuid.UUID `db:"id"`
	IdPengajuan uuid.UUID `db:"id_pengajuan"`
	IdMedis     uuid.UUID `db:"id_barang_medis"`
	Harga       float64   `db:"harga_satuan"`
	Pesanan     int       `db:"jumlah_pesanan"`
	Diterima    int       `db:"jumlah_diterima"`
	Kadaluwarsa time.Time `db:"kadaluwarsa"`
	Batch       string    `db:"no_batch"`
	Updater     uuid.UUID `db:"updater"`
}
