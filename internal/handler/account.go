package handler

import (
	"biling/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var account *models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid JSON provided",
		})
		return
	}

	if account.Id == "" || account.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid account format provided",
		})
		return
	}

	account.Balance = 0
	err := h.Account.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": account.Id})
}

func (h *Handler) GetAccounts(c *gin.Context) {
	accounts, err := h.Account.GetAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get the list of accounts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (h *Handler) GetAccountById(c *gin.Context) {
	id := fmt.Sprint(c.Param("id"))

	account, err := h.Account.GetAccountById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get account by id",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

func (h *Handler) Transaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON provided",
		})
		return
	}

	id, err := h.Account.Transaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unsuccessful operation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
