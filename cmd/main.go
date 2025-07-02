package main

import (
	"stock-app/internal/db"
	"stock-app/internal/handler"
	"stock-app/internal/repository"
	"stock-app/internal/service"
	"stock-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.Init()

	db.AutoMigrate(database)

	stockRepo := repository.NewStockRepository(database)
	stockService := service.NewStockService(stockRepo)
	stockHandler := handler.NewStockHandler(stockService)

	r := gin.Default()
	routes.RegisterRoutes(r, stockHandler)

	r.Run(":8000")
}
