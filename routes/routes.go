package routes

import (
	"Hospital/controllers"
	"Hospital/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	// auth := r.Group("/")
	// auth.Use(middleware.AuthMiddleware(""))
	// {
	// 	auth.GET("/edit/:id", controllers.UpdatePatient)
	// }

	doctor := r.Group("/doctor")
	doctor.Use(middleware.AuthMiddleware("doctor"))
	{
		doctor.GET("/dashboard", controllers.DoctorDashboard)
		doctor.GET("/dashboard/create", controllers.AddPatientView)
		doctor.POST("/dashboard/create", controllers.AddPatient)
		doctor.GET("/dashboard/delete/:id", controllers.DeletePatient)
		doctor.GET("/dashboard/update/:id", controllers.EditPatientView)
		doctor.POST("dashboard/edit/:id", controllers.UpdatePatient)

	}

	receptionist := r.Group("/receptionist")
	receptionist.Use(middleware.AuthMiddleware("receptionist"))
	{
		receptionist.GET("/dashboard", controllers.ReceptionistDashboard)
		receptionist.GET("/dashboard/create", controllers.AddPatientView)
		receptionist.POST("/dashboard/create", controllers.AddPatient)
		receptionist.GET("/dashboard/delete/:id", controllers.DeletePatient)
		receptionist.GET("/dashboard/update/:id", controllers.EditPatientView)
		receptionist.POST("dashboard/edit/:id", controllers.UpdatePatient)

	}
}
