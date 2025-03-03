package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTransactionbyId gets a transaction by id
func (h handler) GetTransactionbyId(c *gin.Context) {
	transactionId := c.Param("transactionId")
	transaction, err := h.TransactionService.GetTransactionbyId(transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

// get last "count" transactions, max is 30
func (h handler) GetTransactionsbyCount(c *gin.Context) {
	accountId := c.Param("account")
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactions, err := h.TransactionService.GetTransactionsbyCount(accountId, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// get transactions in month range
func (h handler) GetTransactionsbyMonthRange(c *gin.Context) {
	
	accountid := c.Param("account")
	startmonth := c.Param("startMonth")
	endmonth := c.Param("endMonth")

	startdate, err := time.Parse("2006-01-02", startmonth+"-01")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enddate, err := time.Parse("2006-01-02", endmonth+"-01")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if enddate.Before(startdate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endmonth cannot be less than startmonth"})
		return
	}

	if enddate.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endmonth cannot be greater than current date"})
		return
	}

	enddate = enddate.AddDate(0, 1, 0)

	transactions, err := h.TransactionService.GetTransactionsbyMonthRange(accountid, startdate, enddate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)

}
