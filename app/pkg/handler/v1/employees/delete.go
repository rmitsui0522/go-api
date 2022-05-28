package employees

import (
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	"github.com/gorilla/mux"
)

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var query model.Employee
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		query.ID = uint(id)

		if err := repo.Delete(&query); err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, map[string]string{
			"message": "",
		})
	}
}
