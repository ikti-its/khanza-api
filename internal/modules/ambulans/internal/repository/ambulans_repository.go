package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/ambulans/internal/entity"
)

type AmbulansRepository interface {
	Insert(ambulans *entity.Ambulans) error
	FindAll() ([]entity.Ambulans, error)
	FindByNoAmbulans(noAmbulans string) (entity.Ambulans, error)
	Update(ambulans *entity.Ambulans) error
	Delete(noAmbulans string) error
	InsertAmbulansRequest(noAmbulans string) error
}
