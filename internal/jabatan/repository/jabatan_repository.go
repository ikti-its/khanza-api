package repository

import "github.com/fathoor/simkes-api/internal/jabatan/entity"

type JabatanRepository interface {
	Insert(jabatan *entity.Jabatan) error
	FindAll() ([]entity.Jabatan, error)
	FindByJabatan(j string) (entity.Jabatan, error)
	Update(jabatan *entity.Jabatan) error
	Delete(jabatan *entity.Jabatan) error
}
