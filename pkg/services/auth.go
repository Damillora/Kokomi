package services

import (
	"errors"
	"time"

	"github.com/Damillora/Kokomi/pkg/config"
	"github.com/Damillora/Kokomi/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) *database.User {
	user := GetUserFromUsername(username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil
	}
	return user
}

func CreateToken(user *database.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Username,
		"iss":   "kokomi-api",
		"sub":   user.ID,
		"aud":   "kokomi",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte(config.CurrentConfig.AuthSecret))
	return jwtToken
}
func CreateRefreshToken(user *database.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"name":  user.Username,
		"iss":   "kokomi-api",
		"sub":   user.ID,
		"aud":   "kokomi-refresh",
		"exp":   time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte(config.CurrentConfig.AuthSecret))
	return jwtToken
}

func ValidateToken(signedToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.CurrentConfig.AuthSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !claims.VerifyExpiresAt(time.Now().Local().Unix(), true) {
		return nil, errors.New("Token is expired")
	}

	return claims, nil
}

func ValidateRefreshToken(signedToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.CurrentConfig.AuthSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !claims.VerifyExpiresAt(time.Now().Local().Unix(), true) {
		return nil, errors.New("Token is expired")
	}

	return claims, nil
}
