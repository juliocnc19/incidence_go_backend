package auth

type DeviceTokenDto struct {
	UserID      uint   `json:"user_id" validate:"required"`
	DeviceToken string `json:"device_token" validate:"required"`
}
