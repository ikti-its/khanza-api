package entity

import "time"

type Tindakan struct {
	NomorRawat   string    `db:"nomor_rawat"`
	NomorRM      string    `db:"nomor_rm"`
	NamaPasien   string    `db:"nama_pasien"`
	Tindakan     *string   `db:"tindakan"` // Use pointers for optional fields
	KodeDokter   *string   `db:"kode_dokter"`
	NamaDokter   *string   `db:"nama_dokter"`
	NIP          *string   `db:"nip"`
	NamaPetugas  *string   `db:"nama_petugas"`
	TanggalRawat time.Time `db:"tanggal_rawat"`
	JamRawat     time.Time `db:"jam_rawat"`
	Biaya        *int64    `db:"biaya"`
}

type JenisTindakan struct {
	Kode         string  `db:"kode" json:"kode"`
	NamaTindakan string  `db:"nama_tindakan" json:"nama_tindakan"`
	Tarif        float64 `db:"total_bayar_dokter_perawat" json:"tarif"`
}
