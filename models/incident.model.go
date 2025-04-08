package models

import (
	"time"
)

type Incident struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string       `gorm:"type:varchar(255);not null" json:"title"`
	Description string       `gorm:"type:text;not null" json:"description"`
	StatusID    uint         `gorm:"not null" json:"status_id"`
	Status      Status       `gorm:"foreignKey:StatusID" json:"status"`
	Response    string       `gorm:"type:text" json:"response"`
	UserID      uint         `gorm:"not null" json:"user_id"`
	User        User         `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	Attachment  []Attachment `json:"attachment"`
}
