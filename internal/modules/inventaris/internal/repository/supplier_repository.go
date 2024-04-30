package repository

import (
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type SupplierRepository interface {
	Find() ([]entity.Supplier, error)
}
