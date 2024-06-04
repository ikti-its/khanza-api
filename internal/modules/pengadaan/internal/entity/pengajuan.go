package entity

import (
	"github.com/google/uuid"
	"time"
)

type Pengajuan struct {
	Id           uuid.UUID `db:"id"`
	Tanggal      time.Time `db:"tanggal_pengajuan"`
	Nomor        string    `db:"nomor_pengajuan"`
	Pegawai      uuid.UUID `db:"id_pegawai"`
	DiskonPersen float64   `db:"diskon_persen"`
	DiskonJumlah float64   `db:"diskon_jumlah"`
	PajakPersen  float64   `db:"pajak_persen"`
	PajakJumlah  float64   `db:"pajak_jumlah"`
	Materai      float64   `db:"materai"`
	Catatan      string    `db:"catatan"`
	Status       string    `db:"status_pesanan"`
	Updater      uuid.UUID `db:"updater"`
}
