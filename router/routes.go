package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/handler"
	"github.com/robert/notification/middleware"
)

func Startserver(r *gin.Engine) {
	r.GET("/", handler.Homepage)
	r.GET("/login", middleware.Isloggedin(), handler.Loginget)
	r.POST("/login", handler.Loginpost)
	r.GET("/secret", middleware.Isloggedin(), handler.Secretpage)
	r.GET("/logout", handler.Logout)
	r.GET("/signup", handler.Signupget)
	r.POST("/signup", handler.Createuser)
	r.Run()
}
