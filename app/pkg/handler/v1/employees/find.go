package employees

import (
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	"github.com/gorilla/mux"
)

func Find() http.HandlerFunc {
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
		employee, err := repo.Find(&query)

		if err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, employee)
	}
}
