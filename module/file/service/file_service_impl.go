package service

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/helper"
	"github.com/fathoor/simkes-api/module/file/model"
	"github.com/fathoor/simkes-api/module/file/validation"
	"path"
)

type fileServiceImpl struct {
}

func (service *fileServiceImpl) Upload(request *model.FileRequest) model.FileResponse {
	filePath, err := validation.ValidateFileRequest(request)
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	fileURL := path.Join("./storage", filePath)

	return model.FileResponse{
		Path: filePath,
		URL:  fileURL,
	}
}

func (service *fileServiceImpl) Get(filetype, filename string) (string, error) {
	return helper.GetFile(filetype, filename)
}

func (service *fileServiceImpl) Delete(filetype, filename string) error {
	filepath, err := helper.GetFile(filetype, filename)
	exception.PanicIfError(err)

	return helper.RemoveFile(filepath)
}

func ProvideFileService() FileService {
	return &fileServiceImpl{}
}
