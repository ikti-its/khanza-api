package entity

import (
	"github.com/google/uuid"
)

type Jadwal struct {
	Id        uuid.UUID `db:"id"`
	IdPegawai uuid.UUID `db:"id_pegawai"`
	IdHari    int       `db:"id_hari"`
	IdShift   string    `db:"id_shift"`
	JamMasuk  string    `db:"jam_masuk"`
	JamPulang string    `db:"jam_pulang"`
	Updater   uuid.UUID `db:"updater"`
}
