package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/golang-jwt/jwt/v5"
)

// SignAccessToken returns a signed access token
func SignAccessToken(user *entity.User) string {
	// const from environment
	secret := os.Getenv("JWT_SECRET")
	accessLifeStr := os.Getenv("JWT_ACCESS_LIFE")

	accessLife, err := strconv.Atoi(accessLifeStr)
	PanicIfError(err)

	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create a new claim
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(accessLife)).Unix()

	// sign the token
	accessToken, err := token.SignedString([]byte(secret))
	PanicIfError(err)

	// return the token
	return accessToken
}

// SignRefreshToken returns a signed refresh token
func SignRefreshToken(user *entity.User) string {
	// const from environment
	secret := os.Getenv("JWT_SECRET")
	refreshLifeStr := os.Getenv("JWT_REFRESH_LIFE")

	refreshLife, err := strconv.Atoi(refreshLifeStr)
	PanicIfError(err)

	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create a new claim
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(refreshLife)).Unix()

	// sign the token
	refreshToken, err := token.SignedString([]byte(secret))
	PanicIfError(err)

	// return the token
	return refreshToken
}
