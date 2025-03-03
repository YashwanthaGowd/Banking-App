package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// use gin framework

// example json request
// {
//   "account_id": "123456",
//   "amount": 100,
//   "type": "credit"
// }

func (h handler) GetTransactionbyId(c *gin.Context) {
	transactionId := c.Param("transactionId") // eg: 123456

	// get transaction
	transaction, err := h.Gateway.GetTransactionbyId(transactionId)
	if err != nil {

		if strings.Contains(err.Error(), "no transactions found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)

}

func (h *handler) GetTransactionsbyMonthRange(c *gin.Context) {
	accountId := c.Param("account")
	startMonth := c.Param("startMonth")
	endMonth := c.Param("endMonth")

	transactions, err := h.Gateway.GetTransactionsbyMonthRange(accountId, startMonth, endMonth)
	if err != nil {

		if strings.Contains(err.Error(), "no transactions found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *handler) GetTransactionsbyAccount(c *gin.Context) {
	account := c.Param("account")
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	transactions, err := h.Gateway.GetTransactionsbyAccount(account, count)
	if err != nil {

		if strings.Contains(err.Error(), "no transactions found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
