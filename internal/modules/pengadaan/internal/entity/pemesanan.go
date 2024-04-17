package entity

import (
	"github.com/google/uuid"
	"time"
)

type Pemesanan struct {
	Id          uuid.UUID `db:"id"`
	Tanggal     time.Time `db:"tanggal_pesan"`
	Nomor       string    `db:"no_pemesanan"`
	IdPengajuan uuid.UUID `db:"id_pengajuan"`
	IdPegawai   uuid.UUID `db:"id_pegawai"`
	Updater     uuid.UUID `db:"updater"`
}
