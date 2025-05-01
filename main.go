package main

import (
	"Hospital/config"
	"Hospital/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"fmt"
	// "Hospital/models"
	// "golang.org/x/crypto/bcrypt"
)

// func seedUser() {
// 	passwordDoctor, _ := bcrypt.GenerateFromPassword([]byte("doctor@123"), bcrypt.DefaultCost)
// 	passwordRecep, _ := bcrypt.GenerateFromPassword([]byte("recep@123"), bcrypt.DefaultCost)

// 	config.DB.Create(&models.User{Username: "doctor", Password: string(passwordDoctor), Role: "doctor"})
// 	config.DB.Create(&models.User{Username: "reception", Password: string(passwordRecep), Role: "receptionist"})
// }

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	config.ConnectDatabase()
	//seedUser()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(r)

	r.Run(":8080")
}
