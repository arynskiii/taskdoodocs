package service

import (
	"archive/zip"
	"doodocs_task/models"
	"github.com/gabriel-vasile/mimetype"
	"mime/multipart"
)

type UnziperService struct{}

func NewUnzipperService() *UnziperService {
	return &UnziperService{}
}

func (s *UnziperService) IsArchive(fileHeader multipart.FileHeader) bool {
	FileType := fileHeader.Header.Get("Content-Type")
	AllowedTypes := []string{"application/x-xz", "application/zip"}
	for _, ttype := range AllowedTypes {
		if ttype != FileType {
			return false
		}

	}
	return true
}

func (s *UnziperService) ProccesingArchive(file multipart.File, fileHeader multipart.FileHeader) (models.Response, error) {
	r, err := zip.NewReader(file, fileHeader.Size)
	if err != nil {
		return models.Response{}, err
	}
	var filesInfo []models.File
	totalsize := float64(0)
	filesCount := float64(0)

	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			reader, err := f.Open()
			if err != nil {
				return models.Response{}, err
			}

			mime, err := mimetype.DetectReader(reader)
			if err != nil {
				return models.Response{}, err
			}

			totalsize += float64(f.FileInfo().Size())
			filesCount++
			FileInfo := models.File{
				File_path: f.Name,
				Size:      float64(f.FileInfo().Size()),
				Mimetype:  mime.String(),
			}
			filesInfo = append(filesInfo, FileInfo)
		}
	}
	sizeOfArchive := float64(fileHeader.Size)
	result := models.Response{
		Filename:     fileHeader.Filename,
		Archive_size: sizeOfArchive,
		Total_size:   totalsize,
		Total_files:  filesCount,
		Files:        filesInfo,
	}

	return result, nil
}
