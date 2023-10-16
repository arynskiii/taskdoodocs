package handler

import (
	"doodocs_task/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("/archive/information", h.unziper)
	api.POST("/archive/files", h.ziper)
	api.POST("/mail/file", h.sendfile)
	return router
}
