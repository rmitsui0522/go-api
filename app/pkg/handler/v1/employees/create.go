package employees

import (
	"encoding/json"
	"net/http"

	"go-api/pkg/model"
	"go-api/pkg/utility"

	validator "github.com/go-playground/validator/v10"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newEmployee model.Employee
		validate := validator.New()

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

		// レコードの重複確認
		if err := esrv.Exists(&newEmployee); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		if err := repo.Create(&newEmployee); err != nil {
			utility.RespondJSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, newEmployee)
	}
}
