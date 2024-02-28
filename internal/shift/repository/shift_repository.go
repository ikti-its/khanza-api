package repository

import "github.com/fathoor/simkes-api/internal/shift/entity"

type ShiftRepository interface {
	Insert(shift *entity.Shift) error
	FindAll() ([]entity.Shift, error)
	FindByNama(n string) (entity.Shift, error)
	Update(shift *entity.Shift) error
	Delete(shift *entity.Shift) error
}
