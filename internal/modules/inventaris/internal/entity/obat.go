package entity

import (
	"github.com/google/uuid"
	"time"
)

type Obat struct {
	Id          uuid.UUID `db:"id"`
	IdMedis     uuid.UUID `db:"id_barang_medis"`
	Industri    int       `db:"id_industri_farmasi"`
	Kandungan   string    `db:"kandungan"`
	SatuanBesar string    `db:"id_satuan_besar"`
	SatuanKecil string    `db:"id_satuan_kecil"`
	Isi         int       `db:"isi"`
	Kapasitas   int       `db:"kapasitas"`
	Jenis       int       `db:"id_jenis"`
	Kategori    int       `db:"id_kategori"`
	Golongan    int       `db:"id_golongan"`
	Kadaluwarsa time.Time `db:"kadaluwarsa"`
	Updater     uuid.UUID `db:"updater"`
}
