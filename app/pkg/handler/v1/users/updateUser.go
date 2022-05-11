package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data model.User
		validate := validator.New()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := validate.Struct(&data); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		user, err := model.UpdateUser(&model.User{ID: uint(id)}, &data)
		if err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}
