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
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, users)
	}
}
