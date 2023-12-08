package service

import (
	"github.com/fathoor/simkes-api/module/file/model"
)

type FileService interface {
	Upload(request *model.FileRequest) model.FileResponse
	Get(filetype, filename string) (string, error)
	Delete(filetype, filename string) error
}
