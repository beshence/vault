package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(g *gin.RouterGroup) {
	g.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "register endpoint not implemented yet",
		})
	})

	g.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "login endpoint not implemented yet",
		})
	})
}
