package controller

import (
	"encoding/json"
	"net/http"

	"go-api/pkg/middleware/auth"
	"go-api/pkg/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"go-api/pkg/utility"
)

type Credentials struct {
	Account  string `json:"mailAddress" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}

func Authentication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Credentials

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := model.FindUser(&model.User{Account: req.Account})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if user.ID == 0 || user.Password != req.Password {
			http.Error(w, "Invalid MailAddress or Password.", http.StatusUnauthorized)
			return
		}

		token, err := createJwtTokenString(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := &AuthenticationResponse{
			Token: token,
		}

		utility.RespondJSON(w, http.StatusOK, res)
	}
}
