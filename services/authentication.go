package services

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func LoginUser(email string, password string) (string, error) {
	user, err := GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	err = compare(user.Password, password)

	if err != nil {
		return "", err
	}

	return createToken(user.ID)
}

func createToken(userId uint) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
