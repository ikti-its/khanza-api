package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/jadwal-pegawai/model"
)

func ValidateJadwalPegawaiRequest(request *model.JadwalPegawaiRequest) error {
	return validation.Validator.Struct(request)
}
