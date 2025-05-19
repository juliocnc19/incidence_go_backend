package dto

type UpdateUserDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	RoleID    uint    `json:"role_id" validate:"required"`
  Cedula    string `json:"cedula" validate:"required"`

}
