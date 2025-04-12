package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Cedula    string    `gorm:"type:varchar(255);unique;not null" json:"cedula"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
