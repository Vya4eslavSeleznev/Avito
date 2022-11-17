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

	if id == -1 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) reserveTransaction(c *gin.Context) {
	var inputData avito.Reservation

	if err := c.BindJSON(&inputData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Transaction.ReserveTransaction(inputData)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getBalance(c *gin.Context) {
	var inputData avito.Balance

	if err := c.BindJSON(&inputData); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	balance, err := h.services.Transaction.GetBalance(inputData)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}
