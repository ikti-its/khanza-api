package service

import "github.com/fathoor/simkes-api/internal/jadwal-pegawai/model"

type JadwalPegawaiService interface {
	Create(request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse
	GetAll() []model.JadwalPegawaiResponse
	GetByNIP(nip string) []model.JadwalPegawaiResponse
	GetByTahunBulan(tahun, bulan int16) []model.JadwalPegawaiResponse
	GetByPK(nip string, tahun, bulan, hari int16) model.JadwalPegawaiResponse
	Update(nip string, tahun, bulan, hari int16, request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse
	Delete(nip string, tahun, bulan, hari int16)
}
