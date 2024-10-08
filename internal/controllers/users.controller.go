package controllers

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/internal/services"
	"github.com/danhbuidcn/go_backend_api/internal/vo"
	"github.com/danhbuidcn/go_backend_api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UsersController {
	return &UsersController{
		userService: userService,
	}
}

func (uc *UsersController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrorCodeParamInvalid, err.Error())
		return
	}
	fmt.Printf("Email param: %s", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
