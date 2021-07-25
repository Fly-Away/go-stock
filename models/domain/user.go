package domain

import "gorm.io/gorm"

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"password"`
	gorm.Model
}