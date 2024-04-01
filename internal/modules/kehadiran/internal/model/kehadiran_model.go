package model

type AttendKehadiranRequest struct {
	IdPegawai       string `json:"id_pegawai" validate:"required,uuid4"`
	IdJadwalPegawai string `json:"id_jadwal_pegawai" validate:"required,uuid4"`
	Tanggal         string `json:"tanggal" validate:"required"`
}

type LeaveKehadiranRequest struct {
	Id        string `json:"id" validate:"required,uuid4"`
	IdPegawai string `json:"id_pegawai" validate:"required,uuid4"`
}

type UpdateKehadiranRequest struct {
	Id              string `json:"id" validate:"required,uuid4"`
	IdPegawai       string `json:"id_pegawai" validate:"required,uuid4"`
	IdJadwalPegawai string `json:"id_jadwal_pegawai" validate:"required,uuid4"`
	Tanggal         string `json:"tanggal" validate:"required"`
	JamMasuk        string `json:"jam_masuk" validate:"required"`
	JamPulang       string `json:"jam_pulang" validate:"required"`
	Keterangan      string `json:"keterangan" validate:"required"`
}

type KehadiranResponse struct {
	Id         string `json:"id"`
	IdPegawai  string `json:"id_pegawai"`
	Tanggal    string `json:"tanggal"`
	JamMasuk   string `json:"jam_masuk"`
	JamPulang  string `json:"jam_pulang,omitempty"`
	Keterangan string `json:"keterangan,omitempty"`
}
