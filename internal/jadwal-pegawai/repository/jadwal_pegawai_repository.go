package repository

import "github.com/fathoor/simkes-api/internal/jadwal-pegawai/entity"

type JadwalPegawaiRepository interface {
	Insert(jadwalPegawai *entity.JadwalPegawai) error
	FindAll() ([]entity.JadwalPegawai, error)
	FindByNIP(nip string) ([]entity.JadwalPegawai, error)
	FindByTahunBulan(tahun, bulan int16) ([]entity.JadwalPegawai, error)
	FindByPK(nip string, tahun, bulan, hari int16) (entity.JadwalPegawai, error)
	Update(jadwalPegawai *entity.JadwalPegawai) error
	Delete(jadwalPegawai *entity.JadwalPegawai) error
}
