package services

import (
	"errors"
	"github.com/Fly-Away/go-stock/database"
	"github.com/Fly-Away/go-stock/models/domain"
	"github.com/Fly-Away/go-stock/models/dto/request/authDto"
	"github.com/Fly-Away/go-stock/models/dto/response"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func Register(user authDto.UserRegisterRequest) (userResp *response.UserResponse, error error, errorCode interface{}) {
	if user.Password != user.ConfirmPassword {
		return &response.UserResponse{}, errors.New("password not match"), 2001
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

	userResponse := response.UserResponse{}
	if err := copier.Copy(&userResponse, &userDomain); err != nil {
		return nil, errors.New("failed to copy User Response"), 2003
	}

	return &userResponse, nil, nil
}

func Login(user authDto.UserLoginRequest) (userResp *response.UserResponse, error error, errorCode interface{}, generatedToken *string) {
	var userDomain domain.User

	database.DB.Where("email = ?", user.Email).First(&userDomain)

	if userDomain.ID == 0 {
		return nil, errors.New("user not found"), 2005, nil
	}

	if err := bcrypt.CompareHashAndPassword(userDomain.Password, []byte(user.Password)); err != nil {
		return nil, errors.New("wrong password"), 2006, nil
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(userDomain.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return nil, errors.New(err.Error()), 2007, nil
	}

	userResponse := response.UserResponse{}
	if errCopy := copier.Copy(&userResponse, &userDomain); errCopy != nil {
		return nil, errors.New(errCopy.Error()), 2008, nil
	}

	return &userResponse, nil, nil, &token
}

type Claims struct {
	jwt.StandardClaims
}

func AuthenticateUser(cookie string) (userResp *response.UserResponse, customError error, errorCode interface{}) {
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return nil, err, 2009
	}

	claims := token.Claims.(*Claims)

	userDomain := domain.User{}

	database.DB.Where("id = ?", claims.Issuer).First(&userDomain)

	userResponse := response.UserResponse{}
	if errCopy := copier.Copy(&userResponse, &userDomain); errCopy != nil {
		return nil, errors.New(errCopy.Error()), 2010
	}

	return &userResponse, nil, nil
}