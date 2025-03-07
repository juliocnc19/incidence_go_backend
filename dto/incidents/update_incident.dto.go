package incidents

type UpdateIncidentDto struct{
	Title          string    `json:"title" validate:"required"`
	Description    string    `json:"description"`
	AttachmentPath string    `json:"attachment_path"`
	StatusID       uint      `json:"status_id" validate:"required"`
	Response       string    `json:"response"`
	UserID         uint      `json:"user_id"`
}
