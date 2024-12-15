package models

import (
	"time"
)

type Ad struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	Title     string `gorm:"not null"`
	Description string `gorm:"not null"`
	CategoryID *uint `gorm:"index"`
	Location string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Images []AdImage `gorm:"constraint:OnDelete:CASCADE"`
}