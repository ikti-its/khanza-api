package validation

import (
	"github.com/fathoor/simkes-api/core/validation"
	"github.com/fathoor/simkes-api/module/auth/model"
)

func ValidateAuthRequest(request *model.AuthRequest) error {
	return validation.Validator.Struct(request)
}
