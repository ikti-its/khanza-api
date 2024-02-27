package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/auth/model"
)

func ValidateAuthRequest(request *model.AuthRequest) error {
	return validation.Validator.Struct(request)
}
