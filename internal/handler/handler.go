package handler

import (
	"biling/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Account service.Account
}

func NewHandler(acc service.Account) *Handler {
	return &Handler{Account: acc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		api.POST("/accounts", h.CreateAccount)
		api.GET("/accounts", h.GetAccounts)
		api.GET("/accounts/:id", h.GetAccountById)
	}

	tr := api.Group("/transaction")
	{
		tr.POST("", h.Transaction)
	}

	return router
}
