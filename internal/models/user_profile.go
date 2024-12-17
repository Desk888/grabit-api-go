package models

import (
	"time"
)

type UserProfile struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          uint   `gorm:"unique;not null"`
	ProfilePicture  string
	PhoneNumber     string
	Address         string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}