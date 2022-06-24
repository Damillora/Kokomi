package models

type LoginFormModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshFormModel struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}