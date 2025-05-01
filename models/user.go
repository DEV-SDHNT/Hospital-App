package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Role     string
}

type PatientDetails struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Name    string
	Illness string
	Address string
	Date    string
	Bed     string
}
