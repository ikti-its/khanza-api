package entity

import (
	"github.com/google/uuid"
	"time"
)

type Bhp struct {
	Id          uuid.UUID `db:"id"`
	IdMedis     uuid.UUID `db:"id_barang_medis"`
	Jumlah      int       `db:"jumlah"`
	Kadaluwarsa time.Time `db:"kadaluwarsa"`
	Updater     uuid.UUID `db:"updater"`
}
