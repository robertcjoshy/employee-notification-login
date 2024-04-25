package main

import (
	"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/handler"
	"github.com/robert/notification/model"
	"github.com/robert/notification/router"
)

func main() {
	model.OpenDatabaseConnection()
	model.AutoMigrateModels()
	r := gin.Default()
	store := handler.Initsession()
	r.Use(sessions.Sessions("mysession", store))
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")
	router.Startserver(r)

}
