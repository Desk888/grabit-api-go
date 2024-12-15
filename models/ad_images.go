package models

import (
	"time"
)

type AdImage struct {
	ID        uint `gorm:"primaryKey"`
	AdID      uint `gorm:"not null"`
	ImageURL string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}