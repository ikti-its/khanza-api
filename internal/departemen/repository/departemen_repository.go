package repository

import "github.com/fathoor/simkes-api/internal/departemen/entity"

type DepartemenRepository interface {
	Insert(departemen *entity.Departemen) error
	FindAll() ([]entity.Departemen, error)
	FindByDepartemen(d string) (entity.Departemen, error)
	Update(departemen *entity.Departemen) error
	Delete(departemen *entity.Departemen) error
}
