package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type SatuanRepository interface {
	Find() ([]entity.Satuan, error)
}
