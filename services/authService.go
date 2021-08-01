package services

import (
	"errors"
	"github.com/Fly-Away/go-stock/database"
	"github.com/Fly-Away/go-stock/models/domain"
	"github.com/Fly-Away/go-stock/util"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func Register(user domain.UserRegisterRequest) (userResp *domain.UserResponse, error error, errorCode interface{}) {
	if user.Password != user.ConfirmPassword {
		return &domain.UserResponse{}, errors.New("password not match"), 2001
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	userDomain := domain.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  encryptedPassword,
	}

	if err := database.DB.Save(&userDomain); err != nil {
		return nil, errors.New(err.Error.Error()), 2002
	}

	userResponse := domain.UserResponse{}
	if err := copier.Copy(&userResponse, &userDomain); err != nil {
		return nil, errors.New("failed to copy User Response"), 2003
	}

	return &userResponse, nil, nil
}

func Login(user domain.UserLoginRequest) (userResp *domain.UserResponse, error error, errorCode interface{}, generatedToken *string) {
	var userDomain domain.User

	database.DB.Where("email = ?", user.Email).First(&userDomain)

	if userDomain.ID == 0 {
		return nil, errors.New("user not found"), 2005, nil
	}

	if err := bcrypt.CompareHashAndPassword(userDomain.Password, []byte(user.Password)); err != nil {
		return nil, errors.New("wrong password"), 2006, nil
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(userDomain.ID)))
	if err != nil {
		return nil, errors.New(err.Error()), 2007, nil
	}

	userResponse := domain.UserResponse{}
	if errCopy := copier.Copy(&userResponse, &userDomain); errCopy != nil {
		return nil, errors.New(errCopy.Error()), 2008, nil
	}

	return &userResponse, nil, nil, &token
}

type Claims struct {
	jwt.StandardClaims
}

func AuthenticateUser(cookie string) (userResp *domain.UserResponse, customError error, errorCode interface{}) {
	id, _ := util.ParseJwt(cookie)

	userDomain := domain.User{}

	database.DB.Where("id = ?", id).First(&userDomain)

	userResponse := domain.UserResponse{}
	if errCopy := copier.Copy(&userResponse, &userDomain); errCopy != nil {
		return nil, errors.New(errCopy.Error()), 2010
	}

	return &userResponse, nil, nil
}