package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Isloggedin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("loggedin"); err == nil {
			c.Next()
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{})
			c.Abort()
		}
	}
}
