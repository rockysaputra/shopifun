package utils

import (
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckLenPassword(password string) bool {
	return len(password) >= 6
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(passwordDB string, passwordInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	return err == nil
}

func GenerateJWTToken(username string, userId uint) (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := jwtToken.SignedString([]byte("Secret"))

	if err != nil {
		return "", err
	}

	return t, nil
}
