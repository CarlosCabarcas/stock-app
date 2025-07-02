package handler

import (
	"net/http"
	"stock-app/internal/service"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	service service.StockService
}

func NewStockHandler(s service.StockService) *StockHandler {
	return &StockHandler{service: s}
}

// GET /stocks
func (h *StockHandler) GetAll(c *gin.Context) {
	stocks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// GET /recommendations
func (h *StockHandler) Recommend(c *gin.Context) {
	stocks, err := h.service.Recommend()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}
