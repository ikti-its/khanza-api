package repository

import "github.com/fathoor/simkes-api/internal/pegawai/entity"

type PegawaiRepository interface {
	Insert(pegawai *entity.Pegawai) error
	FindAll() ([]entity.Pegawai, error)
	FindPage(page, size int) ([]entity.Pegawai, int, error)
	FindByNIP(n string) (entity.Pegawai, error)
	Update(pegawai *entity.Pegawai) error
	Delete(pegawai *entity.Pegawai) error
}
