package routes

import (
	"stock-app/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handler.StockHandler) {
	r.GET("/stocks", h.GetAll)
	r.GET("/recommendations", h.Recommend)
}
