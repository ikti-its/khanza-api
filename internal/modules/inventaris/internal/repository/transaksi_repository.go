package repository

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
)

type TransaksiRepository interface {
	Insert(transaksi *entity.Transaksi) error
	Find() ([]entity.Transaksi, error)
	FindPage(page, size int) ([]entity.Transaksi, int, error)
	FindById(id uuid.UUID) (entity.Transaksi, error)
	Update(transaksi *entity.Transaksi) error
	Delete(transaksi *entity.Transaksi) error
}
