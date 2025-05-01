package controllers

import (
	"Hospital/config"
	"Hospital/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid input"})
		return
	}
	var user models.User

	if err := config.DB.Where("username=?", form.Username).First(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Incorrect Password "})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("role", user.Role)
	session.Save()
	if user.Role == "doctor" {
		print("Doctor found")
		c.Redirect(http.StatusFound, "/doctor/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/receptionist/dashboard")
	}
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
