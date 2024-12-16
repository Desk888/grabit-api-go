package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Profile   UserProfile `gorm:"constraint:OnDelete:CASCADE"`
	Ads       []Ad `gorm:"constraint:OnDelete:CASCADE"`
	Favourites []Favourite `gorm:"constraint:OnDelete:CASCADE"`
	MessageSent []Message `gorm:"foreignKey:SenderID"`
	MessageReceived []Message `gorm:"foreignKey:ReceiverID"`
	ActivityLogs []UserActivity `gorm:"constraint:OnDelete:CASCADE"`
}