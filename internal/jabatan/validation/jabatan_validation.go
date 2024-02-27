package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/jabatan/model"
)

func ValidateJabatanRequest(request *model.JabatanRequest) error {
	return validation.Validator.Struct(request)
}
