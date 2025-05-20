package incidents

type UpdateIncidentDto struct{
	StatusID       uint      `json:"status_id" validate:"required"`
	Response       string    `json:"response"`
}
