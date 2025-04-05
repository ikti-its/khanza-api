package entity

import "time"

type RawatInap struct {
	NomorRawat      string    `json:"nomor_rawat" db:"nomor_rawat"`
	NomorRM         string    `json:"nomor_rm" db:"nomor_rm"`
	NamaPasien      string    `json:"nama_pasien" db:"nama_pasien"`
	AlamatPasien    string    `json:"alamat_pasien" db:"alamat_pasien"`
	PenanggungJawab string    `json:"penanggung_jawab" db:"penanggung_jawab"`
	HubunganPJ      string    `json:"hubungan_pj" db:"hubungan_pj"`
	JenisBayar      string    `json:"jenis_bayar" db:"jenis_bayar"`
	Kamar           string    `json:"kamar" db:"kamar"`
	TarifKamar      float64   `json:"tarif_kamar" db:"tarif_kamar"`
	DiagnosaAwal    string    `json:"diagnosa_awal" db:"diagnosa_awal"`
	DiagnosaAkhir   string    `json:"diagnosa_akhir" db:"diagnosa_akhir"`
	TanggalMasuk    time.Time `json:"tanggal_masuk" db:"tanggal_masuk"`
	JamMasuk        time.Time `json:"jam_masuk" db:"jam_masuk"`
	TanggalKeluar   time.Time `json:"tanggal_keluar" db:"tanggal_keluar"`
	JamKeluar       time.Time `json:"jam_keluar" db:"jam_keluar"`
	TotalBiaya      float64   `json:"total_biaya" db:"total_biaya"`
	StatusPulang    string    `json:"status_pulang" db:"status_pulang"`
	LamaRanap       float64   `json:"lama_ranap" db:"lama_ranap"`
	DokterPJ        string    `json:"dokter_pj" db:"dokter_pj"`
	StatusBayar     string    `json:"status_bayar" db:"status_bayar"`
}
