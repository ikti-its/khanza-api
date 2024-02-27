package repository

import (
	"github.com/fathoor/simkes-api/internal/akun/entity"
)

type AkunRepository interface {
	Insert(akun *entity.Akun) error
	FindAll() ([]entity.Akun, error)
	FindPage(page, size int) ([]entity.Akun, int, error)
	FindByNIP(nip string) (entity.Akun, error)
	Update(akun *entity.Akun) error
	Delete(akun *entity.Akun) error
}
