package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/repository"
	"github.com/jmoiron/sqlx"
)

type brgmedisRepositoryImpl struct {
	DB *sqlx.DB
}

func NewBrgmedisRepository(db *sqlx.DB) repository.BrgmedisRepository {
	return &brgmedisRepositoryImpl{db}
}

func (r *brgmedisRepositoryImpl) Insert(brgmedis *entity.Brgmedis) error {
	query := "INSERT INTO barang_medis VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"

	_, err := r.DB.Exec(query, brgmedis.Id, brgmedis.KodeBarang, brgmedis.IdIndustri, brgmedis.Nama, brgmedis.IdSatBesar, brgmedis.IdSatuan, brgmedis.HDasar, brgmedis.HBeli, brgmedis.HRalan, brgmedis.HKelasI, brgmedis.HKelasII, brgmedis.HKelasIII, brgmedis.HUtama, brgmedis.HVIP, brgmedis.HVVIP, brgmedis.HBeliLuar, brgmedis.HJualBebas, brgmedis.HKaryawan, brgmedis.StokMinimum, brgmedis.IdJenis, brgmedis.Isi, brgmedis.Kapasitas, brgmedis.Kadaluwarsa, brgmedis.IdKategori, brgmedis.IdGolongan)

	return err
}

func (r *brgmedisRepositoryImpl) Find() ([]entity.Brgmedis, error) {
	query := "SELECT id, kode_barang, id_industri, nama, id_satbesar, id_satuan, h_dasar, h_beli, h_ralan, h_kelas1, h_kelas2, h_kelas3, h_utama, h_vip, h_vvip, h_beliluar, h_jualbebas, h_karyawan, stok_minimum, id_jenis, isi, kapasitas, kadaluwarsa, id_kategori, id_golongan FROM barang_medis"

	var records []entity.Brgmedis
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *brgmedisRepositoryImpl) FindById(id uuid.UUID) (entity.Brgmedis, error) {
	query := "SELECT id, kode_barang, id_industri, nama, id_satbesar, id_satuan, h_dasar, h_beli, h_ralan, h_kelas1, h_kelas2, h_kelas3, h_utama, h_vip, h_vvip, h_beliluar, h_jualbebas, h_karyawan, stok_minimum, id_jenis, isi, kapasitas, kadaluwarsa, id_kategori, id_golongan FROM barang_medis WHERE id = $1"

	var record entity.Brgmedis
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *brgmedisRepositoryImpl) Update(brgmedis *entity.Brgmedis) error {
	query := "UPDATE barang_medis SET kode_barang = $2, id_industri = $3, nama = $4, id_satbesar = $5, id_satuan = $6, h_dasar = $7, h_beli = $8, h_ralan = $9, h_kelas1 = $10, h_kelas2 = $11, h_kelas3 = $12, h_utama = $13, h_vip = $14, h_vvip = $15, h_beliluar = $16, h_jualbebas = $17, h_karyawan = $18, stok_minimum = $19, id_jenis = $20, isi = $21, kapasitas = $22, kadaluwarsa = $23, id_kategori = $24, id_golongan = $25 WHERE id = $1"

	_, err := r.DB.Exec(query, brgmedis.Id, brgmedis.KodeBarang, brgmedis.IdIndustri, brgmedis.Nama, brgmedis.IdSatBesar, brgmedis.IdSatuan, brgmedis.HDasar, brgmedis.HBeli, brgmedis.HRalan, brgmedis.HKelasI, brgmedis.HKelasII, brgmedis.HKelasIII, brgmedis.HUtama, brgmedis.HVIP, brgmedis.HVVIP, brgmedis.HBeliLuar, brgmedis.HJualBebas, brgmedis.HKaryawan, brgmedis.StokMinimum, brgmedis.IdJenis, brgmedis.Isi, brgmedis.Kapasitas, brgmedis.Kadaluwarsa, brgmedis.IdKategori, brgmedis.IdGolongan)

	return err
}

func (r *brgmedisRepositoryImpl) Delete(brgmedis *entity.Brgmedis) error {
	query := "DELETE FROM barang_medis WHERE id = $1"

	_, err := r.DB.Exec(query, brgmedis.Id)

	return err
}
