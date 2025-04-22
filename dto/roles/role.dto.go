package roles

type CreateRoleDto struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type UpdateRoleDto struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
} 