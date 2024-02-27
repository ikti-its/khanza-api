package validation

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/role/model"
)

func ValidateRoleRequest(request *model.RoleRequest) error {
	if request.Nama == "Admin" {
		panic(exception.ForbiddenError{
			Message: "You are not allowed to modify this role!",
		})
	}

	return validation.Validator.Struct(request)
}
