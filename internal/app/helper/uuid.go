package helper

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/google/uuid"
)

func MustParse(s string) uuid.UUID {
	id, err := uuid.Parse(s)
	exception.PanicIfError(err, "Failed to parse UUID")

	return id
}
