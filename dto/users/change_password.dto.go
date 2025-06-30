package dto

type ChangePasswordDto struct {
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
