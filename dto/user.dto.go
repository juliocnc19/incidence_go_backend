package dto

type CreateUserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	RoleID    uint   `json:"role_id"`
}
