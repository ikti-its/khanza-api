package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/departemen/model"
)

func ValidateDepartemenRequest(request *model.DepartemenRequest) error {
	return validation.Validator.Struct(request)
}
