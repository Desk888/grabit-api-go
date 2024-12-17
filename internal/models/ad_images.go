package models

import (
	"time"
)

type AdImage struct {
	ID        uint      `gorm:"primaryKey"`
	AdID      uint      `gorm:"not null"`
	ImageURL  string    `gorm:"not null"`
	UploadedAt time.Time `gorm:"autoCreateTime"`
}