package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Berkas struct {
	IdPegawai    uuid.UUID      `gorm:"column:id_pegawai"`
	NIK          string         `gorm:"column:nik"`
	TempatLahir  string         `gorm:"column:tempat_lahir"`
	TanggalLahir time.Time      `gorm:"column:tanggal_lahir"`
	Agama        string         `gorm:"column:agama"`
	Pendidikan   string         `gorm:"column:pendidikan"`
	KTP          string         `gorm:"column:ktp"`
	KK           string         `gorm:"column:kk"`
	NPWP         string         `gorm:"column:npwp"`
	BPJS         string         `gorm:"column:bpjs"`
	Ijazah       string         `gorm:"column:ijazah"`
	SKCK         string         `gorm:"column:skck"`
	STR          string         `gorm:"column:str"`
	SerKom       string         `gorm:"column:serkom"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Updater      uuid.UUID      `gorm:"column:updater"`
}

func (Berkas) TableName() string {
	return "berkas_pegawai"
}
