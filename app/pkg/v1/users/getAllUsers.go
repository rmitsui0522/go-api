package users

import (
	"net/http"

	"go-api/pkg/model"
	"go-api/pkg/utility"
)

func GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := model.FindUsers()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utility.RespondJSON(w, http.StatusOK, users)
	}
}
