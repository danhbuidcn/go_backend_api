package controllers

import (
	"net/http"

	"github.com/danhbuidcn/go_backend_api/internal/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService *services.UserService
}

func NewUsersController() *UsersController {
	return &UsersController{
		userService: services.NewUserService(),
	}
}

// controllers -> services -> repositories -> models -> dbs
func (uc *UsersController) GetUserById(c *gin.Context) {
	uid := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetinfoUserService(),
		"uid":     uid,
		"users":   []string{"cr7", "m10"},
	})
}
