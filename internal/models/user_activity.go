package models

import (
	"time"
)

type UserActivity struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	Action     string    `gorm:"not null"`
	ActionTime time.Time `gorm:"autoCreateTime"`
}