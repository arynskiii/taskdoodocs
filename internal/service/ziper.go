package service

import (
	"archive/zip"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type ZiperrService struct{}

func NewZiperService() *ZiperrService {
	return &ZiperrService{}
}

func (s *ZiperrService) CheckFile(file *multipart.FileHeader) bool {
	t := file.Header.Get("Content-Type")

	AllowedTypes := map[string]bool{
		"image/png":       true,
		"application/xml": true,
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"image/jpeg": true,
		"text/xml":   true, // Добавил от своей руки.
	}

	return AllowedTypes[t]
}

func (s *ZiperrService) CreateZipArchive(files []*multipart.FileHeader) error {
	err := os.Mkdir("./ZIPFILE", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	// Создаю директорию
	zipFile, err := os.Create("./ZIPFILE/urzip.zip") //  Кидаю зипку в только созданную директорию
	if err != nil {

		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err := s.AddToZip(zipWriter, file); err != nil {
			return err
		}

	}

	return nil
}

func (s *ZiperrService) AddToZip(zipWriter *zip.Writer, file *multipart.FileHeader) error {
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Создаем файл в архиве с именем файла
	dstFile, err := zipWriter.Create(file.Filename)
	if err != nil {
		return err
	}

	// Копируем содержимое файла в архив
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
