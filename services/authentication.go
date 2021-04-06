package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthenticationService struct {
	UserService UserService
	EncryptionService EncryptionService
}

func (s *AuthenticationService) LoginUser(email string, password string) (string, error) {
	user, err := s.UserService.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	err = s.EncryptionService.compare(user.Password, password)

	if err != nil {
		return "", err
	}

	return s.createToken(user.ID)
}

func (s *AuthenticationService) ExtractTokenFromHeader(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *AuthenticationService) ExtractTokenMetadata(t string) (uint64, error) {
	token, err := s.verifyToken(t)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, err
	}
	return strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
}

func (s *AuthenticationService)verifyToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *AuthenticationService)createToken(userId uint) (string, error) {
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
