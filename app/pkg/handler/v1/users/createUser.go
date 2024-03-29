package users

import (
	"encoding/json"
	"net/http"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	validator "github.com/go-playground/validator/v10"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		validate := validator.New()

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := validate.Struct(&user); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := model.CreateUser(&user); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}
