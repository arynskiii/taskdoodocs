package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sendfile(c *gin.Context) {
	// Извлекаем список адресов электронной почты
	emails := strings.Split(c.PostForm("emails"), ",")
	if len(emails) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "No recipient emails provided")

		return
	}

	// Извлекаем файл из запроса
	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to retrieve the file")

		return
	}

	// Проверка файла
	if !h.service.Email.CheckFile(fileHeader) {
		newErrorResponse(c, http.StatusBadRequest, "Invalid file type")

		return
	}

	// Преобразуем файл в байты
	fileBytes, err := h.service.Email.FileToBytes(fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to read the file")

		return
	}

	// Отправляем файл на указанные адреса
	err = h.service.Email.SendFileToEmails(fileBytes, fileHeader.Filename, fileHeader.Header.Get("Content-Type"), emails)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to send the file")

		return
	}

	c.JSON(200, gin.H{"message": "File sent successfully"})
}
