package dto

type UpdateUserDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Username  string `json:"username" validate:"required"`
	RoleID    uint    `json:"role_id" validate:"required"`
	AvatarURL string `json:"avatar_url"`
}
