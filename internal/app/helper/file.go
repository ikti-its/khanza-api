package helper

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/google/uuid"
	"os"
	"path"
	"strings"
)

func GenerateFile(fileType string, fileExt string) string {
	fileName := strings.ReplaceAll(uuid.New().String(), "-", "") + fileExt // e.g. e497c0487fde48e5aa7a53f5c2355c67.png
	return path.Join(fileType, fileName)
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func GetFile(fileType string, fileName string) string {
	storage := os.Getenv("APP_STORAGE") // e.g. "/var/www/simkes-api/storage"
	filePath := path.Join(storage, fileType, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) { // Check if file exists
		panic(&exception.NotFoundError{
			Message: fmt.Sprintf("File %s not found", fileName),
		})
	}

	return filePath
}
