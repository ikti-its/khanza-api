package validation

import (
	"fmt"
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/helper"
	"github.com/fathoor/simkes-api/core/validation"
	"github.com/fathoor/simkes-api/module/file/model"
	"path"
)

func ValidateFileRequest(request *model.FileRequest) (string, error) {
	valid := validation.Validator.Struct(request)
	if valid != nil {
		return "", exception.BadRequestError{
			Message: "Invalid request data",
		}
	}

	switch request.Type {
	case "image":
		ext := path.Ext(request.File.Filename)
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			return "", exception.BadRequestError{
				Message: fmt.Sprintf("Invalid file extension: %s is not supported", ext),
			}
		}

		if request.File.Size > 2*1024*1024 {
			return "", exception.BadRequestError{
				Message: "File size exceeds limit",
			}
		}
	default:
		return "", exception.BadRequestError{
			Message: fmt.Sprintf("Invalid file type: %s is not supported", request.Type),
		}
	}

	var (
		fileType = request.Type
		fileName = helper.GenerateFileName()
		fileExt  = path.Ext(request.File.Filename)
		filePath = path.Join(fileType, fileName+fileExt)
	)

	return filePath, nil
}
