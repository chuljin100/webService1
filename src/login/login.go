package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Do(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}
