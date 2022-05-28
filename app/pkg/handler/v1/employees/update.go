package employees

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newEmployee model.Employee
		validate := validator.New()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := validate.Struct(&newEmployee); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		newEmployee.ID = uint(id)

		// レコードの重複確認
		if err := esrv.Exists(&newEmployee); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		employee, err := repo.Updates(&newEmployee)

		if err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, employee)
	}
}
