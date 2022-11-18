package handler

import (
	avito "Avito"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Transaction
// @Description add transaction
// @ID create-transaction
// @Accept  json
// @Produce  json
// @Param input body avito.Accrual true "transaction info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /accrual
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

// @Summary Transaction
// @Description reserve transaction
// @ID reserve-transaction
// @Accept  json
// @Produce  json
// @Param input body avito.Reservation true "reservation info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /reserveTransaction
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

// @Summary Balance
// @Description user balance
// @ID user-balance
// @Accept  json
// @Produce  json
// @Param input body avito.Balance true "balance info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /getBalance
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
