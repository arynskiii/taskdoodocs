package service

import (
	"archive/zip"
	"doodocs_task/internal/config"
	"doodocs_task/models"
	"mime/multipart"
	_ "path/filepath"
)

type Service struct {
	Unziper
	Ziper
	Email
}

type Unziper interface {
	IsArchive(fileHeader multipart.FileHeader) bool
	ProccesingArchive(file multipart.File, fileHeader multipart.FileHeader) (models.Response, error)
}

type Ziper interface {
	CheckFile(file *multipart.FileHeader) bool
	CreateZipArchive(files []*multipart.FileHeader) error
	AddToZip(zipWriter *zip.Writer, file *multipart.FileHeader) error
}

type Email interface {
	FileToBytes(file *multipart.FileHeader) ([]byte, error)
	SendFileToEmails(fileBytes []byte, filename string, mimeType string, emails []string) error
	CheckFile(file *multipart.FileHeader) bool
}

func NewService(cfg config.SMTPConfig) *Service {
	return &Service{
		Unziper: NewUnzipperService(),
		Ziper:   NewZiperService(),
		Email:   NewEmailService(cfg),
	}
}
