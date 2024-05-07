package entity

import (
	"github.com/google/uuid"
	"time"
)

type Darah struct {
	Id          uuid.UUID `db:"id"`
	IdMedis     uuid.UUID `db:"id_barang_medis"`
	Jenis       string    `db:"jenis"`
	Keterangan  string    `db:"keterangan"`
	Kadaluwarsa time.Time `db:"kadaluwarsa"`
	Updater     uuid.UUID `db:"updater"`
}
