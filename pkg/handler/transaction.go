package handler

import (
	avito "Avito"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addTransaction(c *gin.Context) {
	var inputData avito.Accrual

	if err := c.BindJSON(&inputData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Transaction.AddTransaction(inputData)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
