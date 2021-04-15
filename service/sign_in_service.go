package service

import (
	"github.com/gin-gonic/gin"
)

type SignInService struct{}

func (sis *SignInService) CheckIn(c *gin.Context) {
	// bearerToken := c.GetHeader("Authorization")
	// userService := UserService{}
	// jsonResult := userService.getCurrentUserId(bearerToken)
}
