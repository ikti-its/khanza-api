package entity

import (
	"github.com/google/uuid"
	"time"
)

type Mutasi struct {
	Id            uuid.UUID `db:"id"`
	IdBarangMedis uuid.UUID `db:"id_barang_medis"`
	Jumlah        int       `db:"jumlah"`
	Harga         float64   `db:"harga"`
	IdRuanganDari int       `db:"id_ruangandari"`
	IdRuanganKe   int       `db:"id_ruanganke"`
	Tanggal       time.Time `db:"tanggal"`
	Keterangan    string    `db:"keterangan"`
	NoBatch       string    `db:"no_batch"`
	NoFaktur      string    `db:"no_faktur"`
}
