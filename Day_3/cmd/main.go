package main

import (
	"day3/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./static/index.html")
	})
	r.GET("/register", func(c *gin.Context) {
		c.File("./static/register.html")
	})
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.Run(":8080")
}
