package entity

import "github.com/google/uuid"

type Alkes struct {
	Id      uuid.UUID `db:"id"`
	IdMedis uuid.UUID `db:"id_barang_medis"`
	Merek   string    `db:"merek"`
	Updater uuid.UUID `db:"updater"`
}
