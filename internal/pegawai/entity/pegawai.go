package entity

import (
	departemen "github.com/fathoor/simkes-api/internal/departemen/entity"
	jabatan "github.com/fathoor/simkes-api/internal/jabatan/entity"
	"time"
)

type Pegawai struct {
	NIP            string                `gorm:"column:nip;primaryKey"`
	NIK            string                `gorm:"column:nik;not null;unique"`
	Nama           string                `gorm:"column:nama;not null"`
	JenisKelamin   string                `gorm:"column:jenis_kelamin;not null"`
	JabatanNama    string                `gorm:"column:jabatan_nama;not null"`
	Jabatan        jabatan.Jabatan       `gorm:"foreignKey:jabatan_nama;references:nama"`
	DepartemenNama string                `gorm:"column:departemen_nama;not null"`
	Departemen     departemen.Departemen `gorm:"foreignKey:departemen_nama;references:nama"`
	StatusKerja    string                `gorm:"column:status_kerja;not null"`
	Pendidikan     string                `gorm:"column:pendidikan;not null"`
	TempatLahir    string                `gorm:"column:tempat_lahir;not null"`
	TanggalLahir   time.Time             `gorm:"column:tanggal_lahir;not null"`
	Alamat         string                `gorm:"column:alamat;not null"`
	AlamatLat      float64               `gorm:"column:alamat_lat;not null"`
	AlamatLon      float64               `gorm:"column:alamat_lon;not null"`
	Telepon        string                `gorm:"column:telepon;not null"`
	TanggalMasuk   time.Time             `gorm:"column:tanggal_masuk;not null"`
	Foto           string                `gorm:"column:foto;not null"`
	CreatedAt      time.Time             `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time             `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Pegawai) TableName() string {
	return "pegawai"
}
