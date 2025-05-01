package controllers

import (
	"Hospital/config"
	"Hospital/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DoctorDashboard(c *gin.Context) {
	var patients []models.PatientDetails
	config.DB.Find(&patients)
	c.HTML(200, "doctorDashboard.html", gin.H{"patients": patients})
}
func ReceptionistDashboard(c *gin.Context) {
	var patients []models.PatientDetails
	config.DB.Find(&patients)
	c.HTML(200, "receptionistDashboard.html", gin.H{"patients": patients})
}
func AddPatientView(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	c.HTML(http.StatusOK, "addPatient.html", gin.H{"role": role})
}

func AddPatient(c *gin.Context) {
	patient := models.PatientDetails{
		Name:    c.PostForm("Name"),
		Illness: c.PostForm("Illness"),
		Address: c.PostForm("Address"),
		Date:    c.PostForm("Date"),
		Bed:     c.PostForm("Bed"),
	}
	config.DB.Create(&patient)
	session := sessions.Default(c)
	role := session.Get("role")
	if role == "receptionist" {
		c.Redirect(http.StatusFound, "/receptionist/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/doctor/dashboard")
	}
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.PatientDetails{}, id)
	session := sessions.Default(c)
	role := session.Get("role")
	fmt.Print(role)
	if role == "receptionist" {
		c.Redirect(http.StatusFound, "/receptionist/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/doctor/dashboard")
	}
}

func EditPatientView(c *gin.Context) {
	var patient models.PatientDetails
	session := sessions.Default(c)
	role := session.Get("role")
	config.DB.First(&patient, c.Param("id"))
	c.HTML(http.StatusOK, "editPatient.html", gin.H{"patient": patient, "role": role})
}

func UpdatePatient(c *gin.Context) {
	var patient models.PatientDetails
	config.DB.First(&patient, c.Param("id"))

	patient.Name = c.PostForm("Name")
	patient.Illness = c.PostForm("Illness")
	patient.Address = c.PostForm("Address")
	patient.Date = c.PostForm("Date")
	patient.Bed = c.PostForm("Bed")

	config.DB.Save(&patient)
	session := sessions.Default(c)
	role := session.Get("role")
	fmt.Println(role, "--------------------")
	if role == "receptionist" {
		c.Redirect(http.StatusFound, "/receptionist/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/doctor/dashboard")
	}
}
