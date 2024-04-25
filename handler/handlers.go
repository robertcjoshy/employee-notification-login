package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/robert/notification/model"
)

func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Loginget(c *gin.Context) {
	c.HTML(http.StatusOK, "secret.html", gin.H{})
}

func Loginpost(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := model.Authenticate(username, c); err == nil {
		c.SetCookie("loggedin", "true", 3600, "", "", false, true)
		//c.HTML(http.StatusOK, "secret.html", gin.H{})
		c.Redirect(http.StatusFound, "/secret")

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"user": username,
			"pass": password,
		})
	}
}

func Secretpage(c *gin.Context) {
	// if _, err := c.Cookie("loggedin"); err == nil {
	// 	c.HTML(http.StatusOK, "secret.html", gin.H{})
	// } else {
	// 	c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
	// }
	c.HTML(http.StatusFound, "secret.html", gin.H{})
}

func Logout(c *gin.Context) {
	c.SetCookie("loggedin", "", -1, "", "", false, true)
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Signupget(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func Createuser(c *gin.Context) {
	var input model.User
	if err := c.Bind(&input); err != nil { // Bind form data to the user struct
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind json"})
		return
	}

	err := input.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.SetCookie("loggedin", "true", 3600, "", "", false, true)
	c.HTML(http.StatusFound, "secret.html", gin.H{})
}
