package incidents

type CreateIncidentDto struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	StatusID     uint   `json:"status_id" validate:"required"`
	Response     string `json:"response"`
	UserID       uint   `json:"user_id"`
}
