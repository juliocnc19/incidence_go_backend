package models

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Title          string    `gorm:"type:varchar(255);not null"`
	CategoryID     uint      `gorm:"not null"`
	Category       Category  `gorm:"foreignKey:CategoryID"`
	Description    string    `gorm:"type:text;not null"`
	AttachmentPath string    `gorm:"type:varchar(255)"`
	StatusID       uint      `gorm:"not null"`
	Status         Status    `gorm:"foreignKey:StatusID"`
	Response       string    `gorm:"type:text"`
	UserID         uint      `gorm:"not null"`
	User           User      `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
