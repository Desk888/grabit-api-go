package models

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`
	Ads  []Ad   `gorm:"many2many:ad_tags"`
}