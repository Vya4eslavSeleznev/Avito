package handler

import (
	"Avito/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/accrual", h.addTransaction)
	router.POST("/reserveTransaction", h.reserveTransaction)

	return router
}
