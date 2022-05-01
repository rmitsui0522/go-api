package controller

import (
	"fmt"
	"net/http"
	"time"

	"go-api/pkg/middleware/auth"
	"go-api/pkg/model"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

type Credentials struct {
	Account  string `json:"mailAddress" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.String(http.StatusOK, "Welcome "+name+"!")
	}
}

func Authentication() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req Credentials
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		user, err := model.FindUser(&model.User{Account: req.Account})
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		if user.ID == 0 || user.Password != req.Password {
			return c.String(http.StatusUnauthorized, "Invalid MailAddress or Password.")
		}

		token, err := createJwtTokenString(user)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		res := &AuthenticationResponse{
			Token: token,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func Authorization(c echo.Context) (jwt.Claims, error) {
	var req AuthenticationResponse
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	claims, err := validateToken(req.Token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func createJwtTokenString(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"Name":      user.FirstName + user.LastName,
		"Account":   user.Account,
		"Admin":     user.ID == 1,
		"ExpiresAt": time.Now().Add(time.Hour * 72).Unix(),
		"Id":        fmt.Sprint(user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(auth.HmacSecret())

	return tokenString, err
}

func validateToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return auth.HmacSecret(), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
