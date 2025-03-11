package models

import (
	"time"
)

type UserToken struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	DeviceToken string    `gorm:"type:varchar(255);unique;not null" json:"device_token"`
	Platform    string    `gorm:"type:varchar(50);not null" json:"plataform"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
