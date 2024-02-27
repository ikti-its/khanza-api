package helper

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/google/uuid"
	"os"
	"path"
)

func GenerateFile(fileType string, fileExt string) string {
	return path.Join(fileType, uuid.New().String()+fileExt)
}

func RemoveFile(filepath string) error {
	return os.Remove(filepath)
}

func GetFile(filetype string, filename string) (string, error) {
	cfg := config.ProvideConfig()
	storage := cfg.Get("APP_STORAGE")
	filePath := path.Join(storage, filetype, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", exception.NotFoundError{
			Message: "File not found",
		}
	}

	return filePath, nil
}
