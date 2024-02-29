package validation

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/kehadiran/model"
	"time"
)

func ValidateKehadiranRequest(request *model.KehadiranRequest) error {
	if _, err := time.Parse("2006-01-02", request.Tanggal); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	return validation.Validator.Struct(request)
}

func ValidateKehadiranUpdateRequest(request *model.KehadiranUpdateRequest) error {
	if _, err := time.Parse("2006-01-02", request.Tanggal); err != nil {
		return exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	if request.JamMasuk != "" {
		if _, err := time.Parse("15:04:05", request.JamMasuk); err != nil {
			return exception.BadRequestError{
				Message: "Invalid time format",
			}
		}
	}

	if request.JamKeluar != "" {
		if _, err := time.Parse("15:04:05", request.JamKeluar); err != nil {
			return exception.BadRequestError{
				Message: "Invalid time format",
			}
		}
	}

	return validation.Validator.Struct(request)
}
