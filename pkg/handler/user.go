package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUsers(c *gin.Context) {

}

func (h *Handler) addUser(c *gin.Context) {
	id, err := h.services.AddUser()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
