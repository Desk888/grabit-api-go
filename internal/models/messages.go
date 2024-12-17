package models

import (
	"time"
)

type Message struct {
	ID            uint      `gorm:"primaryKey"`
	SenderID      uint      `gorm:"not null"`
	ReceiverID    uint      `gorm:"not null"`
	AdID          *uint
	MessageContent string   `gorm:"not null"`
	SentAt        time.Time `gorm:"autoCreateTime"`
}