package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/organisasi/internal/entity"
)

type OrganisasiRepository interface {
	Insert(organisasi *entity.Organisasi) error
	Find() ([]entity.Organisasi, error)
	FindPage(page, size int) ([]entity.Organisasi, int, error)
	FindCurrent() (entity.Organisasi, error)
	FindById(id uuid.UUID) (entity.Organisasi, error)
	Update(organisasi *entity.Organisasi) error
	Delete(organisasi *entity.Organisasi) error
}
