package helper

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/app/exception"
)

func MustParse(s string) uuid.UUID {
	id, err := uuid.Parse(s)
	exception.PanicIfError(err, "Failed to parse UUID")

	return id
}
