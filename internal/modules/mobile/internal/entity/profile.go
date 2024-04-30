package entity

import "github.com/google/uuid"

type Profile struct {
	Akun      uuid.UUID `db:"akun"`
	Foto      string    `db:"foto"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Alamat    string    `db:"alamat"`
	AlamatLat float64   `db:"alamat_lat"`
	AlamatLon float64   `db:"alamat_lon"`
	Kota      string    `db:"kota"`
	KodePos   string    `db:"kode_pos"`
	Updater   uuid.UUID `db:"updater"`
}
