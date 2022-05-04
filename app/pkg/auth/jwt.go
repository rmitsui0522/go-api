package auth

import (
	"fmt"
	"time"

	"go-api/pkg/model"

	"github.com/form3tech-oss/jwt-go"
)

func NewJwtTokenString(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"Name":      user.FirstName + user.LastName,
		"Account":   user.Account,
		"Admin":     user.ID == 1,
		"ExpiresAt": time.Now().Add(time.Hour * 72).Unix(),
		"Id":        fmt.Sprint(user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(HmacSecret())

	return tokenString, err
}

func ValidateToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return HmacSecret(), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
