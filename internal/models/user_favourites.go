package models

import (
	"time"
)

type Favourite struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	AdID      uint `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}