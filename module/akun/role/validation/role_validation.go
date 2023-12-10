package validation

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/validation"
	"github.com/fathoor/simkes-api/module/akun/role/model"
)

func ValidateRoleRequest(request *model.RoleRequest) error {
	if request.Name == "Admin" {
		panic(exception.ForbiddenError{
			Message: "Forbidden role name",
		})
	}
	return validation.Validator.Struct(request)
}
