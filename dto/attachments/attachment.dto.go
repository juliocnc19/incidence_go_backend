package attachments

type CreateAttachmentDto struct {
	AttachmentPath string `json:"attachment_path" validate:"required"`
	IncidentID     uint   `json:"incident_id" validate:"required"`
}

type UpdateAttachmentDto struct {
	AttachmentPath string `json:"attachment_path" validate:"required"`
	IncidentID     uint   `json:"incident_id" validate:"required"`
} 