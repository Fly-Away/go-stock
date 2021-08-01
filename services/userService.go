package services

import (
	"github.com/Fly-Away/go-stock/database"
	"github.com/Fly-Away/go-stock/models/domain"
)

func GetAllUsers() ([]domain.User, error) {
	var userDomain []domain.User
	database.DB.Find(&userDomain)

	return userDomain, nil
}