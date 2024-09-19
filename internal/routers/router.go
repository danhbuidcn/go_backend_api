package routers

import (
	"fmt"

	c "github.com/danhbuidcn/go_backend_api/internal/controllers"
	"github.com/danhbuidcn/go_backend_api/internal/middlerwares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ->> AA")
		c.Next()
		fmt.Println("Alter ->> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ->> BB")
		c.Next()
		fmt.Println("Alter ->> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before ->> CC")
	c.Next()
	fmt.Println("Alter ->> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	// use the middlerware
	// r.Use(AA(), BB(), CC)
	r.Use(middlerwares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("v1")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)
		v1.GET("/user/:id", c.NewUsersController().GetUserById)
	}
	return r
}
