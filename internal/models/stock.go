package models

import (
	"time"
)

type Stock struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Symbol    string  `gorm:"type:varchar(10);not null;unique"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
	Currency  string  `gorm:"type:varchar(5);default:'USD'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
