package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/pegawai/model"
)

func ValidatePegawaiRequest(request *model.PegawaiRequest) error {
	return validation.Validator.Struct(request)
}
