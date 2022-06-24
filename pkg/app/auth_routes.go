package app

import (
	"net/http"

	"github.com/Damillora/Kokomi/pkg/models"
	"github.com/Damillora/Kokomi/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializeAuthRoutes(g *gin.Engine) {
	g.POST("/api/auth/login", createToken)
	g.POST("/api/auth/refresh", refreshToken)
}
func createToken(c *gin.Context) {
	var model models.LoginFormModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	user := services.Login(model.Username, model.Password)

	if user != nil {
		token := services.CreateToken(user)
		refreshToken := services.CreateRefreshToken(user)
		c.JSON(http.StatusOK, models.TokenResponse{
			Token: token,
			RefreshToken: refreshToken,
		})

	} else {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Wrong username or password",
		})
	}
}
func refreshToken(c *gin.Context) {
	var model models.RefreshFormModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}
    claims, err := services.ValidateRefreshToken(model.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	user := services.GetUser(claims["sub"].(string))

	if user != nil {
		token := services.CreateToken(user)
		refreshToken := services.CreateRefreshToken(user)
		c.JSON(http.StatusOK, models.TokenResponse{
			Token: token,
			RefreshToken: refreshToken,
		})

	} else {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Refresh token wrong",
		})
	}
}