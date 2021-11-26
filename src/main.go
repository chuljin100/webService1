package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"login"
	"signup"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/signup", signup.Do)
		v1.GET("/login", login.Do)
	}
	router.Run(":8080")
}
