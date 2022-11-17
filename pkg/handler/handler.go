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

	router.GET("/get_users", h.getUsers)
	router.POST("/add_user", h.addUser)

	router.POST("/add_transaction", h.addTransaction)

	return router
}
