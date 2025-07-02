package db

import (
	"log"
	"stock-app/internal/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Stock{},
	)

	if err != nil {
		log.Fatalf("❌ Error en la migración: %v", err)
	} else {
		log.Println("✅ Migraciones completadas")
	}
}
