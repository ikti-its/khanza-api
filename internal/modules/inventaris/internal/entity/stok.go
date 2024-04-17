package entity

import (
	"github.com/google/uuid"
	"time"
)

type Stok struct {
	Id         uuid.UUID `db:"id"`
	Nomor      string    `db:"no_keluar"`
	IdMedis    uuid.UUID `db:"id_barang_medis"`
	IdPegawai  uuid.UUID `db:"id_pegawai"`
	Tanggal    time.Time `db:"tanggal_stok_keluar"`
	Jumlah     int       `db:"jumlah_keluar"`
	Keterangan string    `db:"keterangan"`
	Updater    uuid.UUID `db:"updater"`
}
