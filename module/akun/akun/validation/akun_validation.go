package validation

import (
	"github.com/fathoor/simkes-api/core/validation"
	"github.com/fathoor/simkes-api/module/akun/akun/model"
)

func ValidateAkunRequest(request *model.AkunRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateAkunUpdateRequest(request *model.AkunUpdateRequest) error {
	return validation.Validator.Struct(request)
}
