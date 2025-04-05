package models

import "time"

type Attachment struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AttachmentPath string    `gorm:"type:varchar(255)" json:"attachment_path"`
	IncidentID     uint      `json:"incident_id"`
	Incident       Incident  `gorm:"foreignKey:IncidentID" json:"Incident"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
