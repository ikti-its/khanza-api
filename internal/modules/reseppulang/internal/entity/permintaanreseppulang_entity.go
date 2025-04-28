package entity

import "time"

// PermintaanResepPulang maps to the permintaan_resep_pulang table
type PermintaanResepPulang struct {
	NoPermintaan  string     `db:"no_permintaan"`  // varchar(14), PK
	TglPermintaan *time.Time `db:"tgl_permintaan"` // nullable DATE
	Jam           time.Time  `db:"jam"`            // TIME
	NoRawat       string     `db:"no_rawat"`       // varchar(17)
	KdDokter      string     `db:"kd_dokter"`      // varchar(20)
	Status        string     `db:"status"`         // ENUM-like: 'Sudah', 'Belum'
	TglValidasi   time.Time  `db:"tgl_validasi"`   // DATE
	JamValidasi   time.Time  `db:"jam_validasi"`   // TIME
}
