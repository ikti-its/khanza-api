package validation

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/cuti/model"
	"time"
)

func ValidateCutiCreateRequest(request *model.CutiCreateRequest) error {
	if _, err := time.Parse("2006-01-02", request.TanggalMulai); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	if _, err := time.Parse("2006-01-02", request.TanggalSelesai); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	return validation.Validator.Struct(request)
}

func ValidateCutiUpdateRequest(request *model.CutiUpdateRequest) error {
	if _, err := time.Parse("2006-01-02", request.TanggalMulai); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	if _, err := time.Parse("2006-01-02", request.TanggalSelesai); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	return validation.Validator.Struct(request)
}
