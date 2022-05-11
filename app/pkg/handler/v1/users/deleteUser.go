package users

import (
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	"github.com/gorilla/mux"
)

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		user, err := model.DeleteUser(&model.User{ID: uint(id)})
		if err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, user)
	}
}
