package helper

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/google/uuid"
	"os"
	"path"
)

func GenerateFileName() string {
	return uuid.New().String()
}

func RemoveFile(filepath string) error {
	return os.Remove(filepath)
}

func GetFile(filetype string, filename string) (string, error) {
	resource := "./resource/"
	filePath := path.Join(resource, filetype, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", exception.NotFoundError{
			Message: "File not found",
		}
	}

	return filePath, nil
}
