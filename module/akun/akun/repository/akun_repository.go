package repository

import "github.com/fathoor/simkes-api/module/akun/akun/entity"

type AkunRepository interface {
	Insert(akun *entity.Akun) error
	FindAll() ([]entity.Akun, error)
	FindByNIP(nip string) (entity.Akun, error)
	Update(akun *entity.Akun) error
	Delete(akun *entity.Akun) error
}
