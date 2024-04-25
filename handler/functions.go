package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Authenticated(user, pass string) bool {
	if user == "robert" && pass == "1234" {
		return true
	}
	return false
}

func Initsession() sessions.Store {
	store := cookie.NewStore([]byte("secret_key"))
	return store
}

func Setsessions(c *gin.Context, key, value string) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

func clearsession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
