package service

import (
	"doodocs_task/models"
	"doodocs_task/util"
	"testing"
)

func CreateRandomArchiveFile(t *testing.T) models.Response {
	var files []models.File

	file := models.File{
		File_path: util.RandomFilePath(),
		Size:      util.RandomSize(),
		Mimetype:  util.RandomMimeType(),
	}

	files = append(files, file)
	arg := models.Response{
		Filename:     util.RandomFileName(),
		Archive_size: util.RandomArchiveSize(),
		Total_size:   util.RandomTotalSize(),
		Total_files:  util.RandomTotalFiles(),
		Files:        files,
	}
	return arg
}
