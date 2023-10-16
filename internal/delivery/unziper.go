package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxMemory = 10 << 20 // 1MB

func (h *Handler) unziper(c *gin.Context) {
	err := c.Request.ParseMultipartForm(maxMemory)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to parse multipart form")

		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to get file from form")

		return
	}

	if h.service.Unziper.IsArchive(*fileHeader) {
		newErrorResponse(c, http.StatusBadRequest, "File is not a valid archive")

		return
	}

	// Обработка архива и извлечение информации о файлах
	aboutFile, err := h.service.Unziper.ProccesingArchive(file, *fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to process the archive")

		return
	}

	// Возврат информации о файлах в архиве
	c.JSON(http.StatusOK, aboutFile)
}
