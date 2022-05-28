package employees

import (
	"net/http"

	"go-api/pkg/utility"
)

func FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := repo.FindAll()

		if err != nil {
			utility.RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
			return
		}

		utility.RespondJSON(w, http.StatusOK, employees)
	}
}
