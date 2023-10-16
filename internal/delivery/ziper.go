package handler

import (
	"archive/zip"
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ziper(c *gin.Context) {
	// Парсим многопоточные данные
	err := c.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to parse multipart form")
		return
	}
	files := c.Request.MultipartForm.File["files[]"]
	for _, file := range files {
		if !h.service.Ziper.CheckFile(file) {

			newErrorResponse(c, http.StatusBadRequest, "cannot get this type of file")
			return
		}
	}

	// Создаем буфер для хранения архива
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// Добавляем файлы в архив
	for _, file := range files {
		if err := h.service.Ziper.AddToZip(zipWriter, file); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "Failed to add file to ZIP archive")
			return
		}
	}
	if err := h.service.Ziper.CreateZipArchive(files); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to create ZIP archive")

		return
	}

	// Закрываем архив
	err = zipWriter.Close()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to close ZIP archive")
		return
	}

	// Отправляем бинарные данные архива в теле ответа
	c.Data(http.StatusOK, "application/zip", buf.Bytes())
}
