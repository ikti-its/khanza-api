package entity

type Role struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type Jabatan struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type Departemen struct {
	Id   int    `db:"id"`
	Nama string `db:"nama"`
}

type StatusAktif struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}

type Shift struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}

type AlasanCuti struct {
	Id   string `db:"id"`
	Nama string `db:"nama"`
}
