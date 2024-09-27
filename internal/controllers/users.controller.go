package controllers

import (
	"github.com/danhbuidcn/go_backend_api/internal/services"
	"github.com/danhbuidcn/go_backend_api/pkg/response"
	"github.com/gin-gonic/gin"
)

// type UsersController struct {
// 	userService *services.UserService
// }

// func NewUsersController() *UsersController {
// 	return &UsersController{
// 		userService: services.NewUserService(),
// 	}
// }

// // controllers -> services -> repositories -> models -> dbs
// func (uc *UsersController) GetUserById(c *gin.Context) {
// 	response.SuccessResponse(c, 2001, []string{"tipjs", "m10"})
// 	// response.ErrorResponse(c, 2003, "error message")
// }

// INTERFACE VERSION

type UsersController struct {
	userService services.IUserService
}

func NewUsersController(userService services.IUserService) *UsersController {
	return &UsersController{
		userService: userService,
	}
}

func (uc *UsersController) Register(c *gin.Context) {
	result := uc.userService.Resgister("", "")
	response.SuccessResponse(c, result, nil)
}
