package model

import "mime/multipart"

type FileRequest struct {
	File *multipart.FileHeader `json:"file" validate:"required"`
	Type string                `json:"type" validate:"required,alpha"`
}

type FileResponse struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
