package entity

import "time"

type PemberianObat struct {
	TanggalBeri time.Time `db:"tanggal_beri"`
	JamBeri     time.Time `db:"jam_beri"`
	NomorRawat  string    `db:"nomor_rawat"`
	NamaPasien  string    `db:"nama_pasien"`
	KodeObat    string    `db:"kode_obat"`
	NamaObat    string    `db:"nama_obat"`
	Embalase    *string   `db:"embalase"`
	Tuslah      *string   `db:"tuslah"`
	Jumlah      *string   `db:"jumlah"`
	BiayaObat   *float64  `db:"biaya_obat"`
	Total       *float64  `db:"total"`
	Gudang      *string   `db:"gudang"`
	NoBatch     *string   `db:"no_batch"`
	NoFaktur    *string   `db:"no_faktur"`
}
