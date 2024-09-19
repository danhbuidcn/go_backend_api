package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("---> My Handler")

	name := c.Param("name")
	uid := c.DefaultQuery("uid", "uid")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10"},
	})
}
