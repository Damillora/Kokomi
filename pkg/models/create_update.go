package models

type UserCreateModel struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateModel struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
