package statuses

type CreateStatusDto struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type UpdateStatusDto struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
} 