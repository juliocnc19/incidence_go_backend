package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	DeviceToken string    `gorm:"type:varchar(255);unique;not null"`
	Platform    string    `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
