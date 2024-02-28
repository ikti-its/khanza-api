package model

type PegawaiRequest struct {
	NIP            string  `json:"nip" validate:"required,alphanum,max=5"`
	NIK            string  `json:"nik" validate:"required,numeric,max=16"`
	Nama           string  `json:"nama" validate:"required,alpha,max=50"`
	JenisKelamin   string  `json:"jenis_kelamin" validate:"required,oneof=L P"`
	JabatanNama    string  `json:"jabatan_nama" validate:"required,alpha,max=25"`
	DepartemenNama string  `json:"departemen_nama" validate:"required,alpha,max=25"`
	StatusKerja    string  `json:"status_kerja" validate:"required,oneof=Tetap Kontrak"`
	Pendidikan     string  `json:"pendidikan" validate:"required,oneof=SD SMP SMA D3 S1 S2 S3"`
	TempatLahir    string  `json:"tempat_lahir" validate:"required,alpha,max=20"`
	TanggalLahir   string  `json:"tanggal_lahir" validate:"required"`
	Alamat         string  `json:"alamat" validate:"required"`
	AlamatLat      float64 `json:"alamat_lat" validate:"omitempty"`
	AlamatLon      float64 `json:"alamat_lon" validate:"omitempty"`
	Telepon        string  `json:"telepon" validate:"required,numeric,max=15"`
	TanggalMasuk   string  `json:"tanggal_masuk" validate:"required"`
	Foto           string  `json:"foto" validate:"omitempty"`
}

type PegawaiResponse struct {
	NIP            string  `json:"nip"`
	NIK            string  `json:"nik"`
	Nama           string  `json:"nama"`
	JenisKelamin   string  `json:"jenis_kelamin"`
	JabatanNama    string  `json:"jabatan_nama"`
	DepartemenNama string  `json:"departemen_nama"`
	StatusKerja    string  `json:"status_kerja"`
	Pendidikan     string  `json:"pendidikan"`
	TempatLahir    string  `json:"tempat_lahir"`
	TanggalLahir   string  `json:"tanggal_lahir"`
	Alamat         string  `json:"alamat"`
	AlamatLat      float64 `json:"alamat_lat"`
	AlamatLon      float64 `json:"alamat_lon"`
	Telepon        string  `json:"telepon"`
	TanggalMasuk   string  `json:"tanggal_masuk"`
	Foto           string  `json:"foto"`
}

type PegawaiPageResponse struct {
	Pegawai []PegawaiResponse `json:"pegawai"`
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Total   int               `json:"total"`
}
