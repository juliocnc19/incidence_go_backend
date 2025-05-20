package dto

type CreateUserDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Username  string `json:"username" validate:"required"`
	RoleID    uint   `json:"role_id" validate:"required"`
	Cedula    string `json:"cedula" validate:"required"`
}
