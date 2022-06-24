package models

type TokenResponse struct {
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserProfileResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}