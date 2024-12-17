package models

import (
	"time"
)

type Favorite struct {
	ID      uint      `gorm:"primaryKey"`
	UserID  uint      `gorm:"not null"`
	AdID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}