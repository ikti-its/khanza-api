package validation

import (
	"github.com/fathoor/simkes-api/core/validation"
	"github.com/fathoor/simkes-api/module/akun/role/model"
)

func ValidateRoleRequest(request *model.RoleRequest) error {
	return validation.Validator.Struct(request)
}
