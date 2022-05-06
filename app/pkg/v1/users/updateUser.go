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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := model.UpdateUser(&model.User{ID: uint(id)}, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}
