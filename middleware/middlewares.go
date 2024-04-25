package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
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
*/
func Isloggedin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.Next()
		}

	}
}
