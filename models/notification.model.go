package models

import (
	"time"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;" gorm:"foreignKey:UserID" json:"user"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
