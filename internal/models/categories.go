package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Ads  []Ad   `gorm:"constraint:OnDelete:SET NULL"`
}