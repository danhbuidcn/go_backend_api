package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/ping/:name", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.DefaultQuery("uid", "uid")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10"},
	})
}
