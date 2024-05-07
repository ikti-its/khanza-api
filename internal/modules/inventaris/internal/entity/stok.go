package entity

import (
	"github.com/google/uuid"
	"time"
)

type Stok struct {
	Id          uuid.UUID `db:"id"`
	IdTransaksi uuid.UUID `db:"id_transaksi_keluar_barang_medis"`
	Nomor       string    `db:"no_keluar"`
	IdPegawai   uuid.UUID `db:"id_pegawai"`
	Tanggal     time.Time `db:"tanggal_stok_keluar"`
	Keterangan  string    `db:"keterangan"`
	Updater     uuid.UUID `db:"updater"`
}
