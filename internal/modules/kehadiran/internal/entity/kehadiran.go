package entity

import (
	"github.com/google/uuid"
	"time"
)

type Kehadiran struct {
	Id              uuid.UUID `db:"id"`
	IdPegawai       uuid.UUID `db:"id_pegawai"`
	IdJadwalPegawai uuid.UUID `db:"id_jadwal_pegawai"`
	Tanggal         time.Time `db:"tanggal"`
	JamMasuk        time.Time `db:"jam_masuk"`
	JamPulang       time.Time `db:"jam_pulang"`
	Keterangan      string    `db:"keterangan"`
	Updater         uuid.UUID `db:"updater"`
}
