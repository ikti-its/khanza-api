package validation

import (
	"github.com/fathoor/simkes-api/internal/akun/model"
	"github.com/fathoor/simkes-api/internal/app/validation"
)

func ValidateAkunRequest(request *model.AkunRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateAkunUpdateRequest(request *model.AkunUpdateRequest) error {
	return validation.Validator.Struct(request)
}
