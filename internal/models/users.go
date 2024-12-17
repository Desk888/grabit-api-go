package models

import (
	"time"
)

type User struct {
	ID           uint           `gorm:"primaryKey"`
	Username     string         `gorm:"uniqueIndex;not null"`
	Email        string         `gorm:"uniqueIndex;not null"`
	PasswordHash string         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	Profile      UserProfile    `gorm:"constraint:OnDelete:CASCADE"`
	Ads          []Ad           `gorm:"constraint:OnDelete:CASCADE"`
	Favorites    []Favorite     `gorm:"constraint:OnDelete:CASCADE"`
	MessagesSent []Message      `gorm:"foreignKey:SenderID"`
	MessagesRecv []Message      `gorm:"foreignKey:ReceiverID"`
	ActivityLogs []UserActivity `gorm:"constraint:OnDelete:CASCADE"`
}