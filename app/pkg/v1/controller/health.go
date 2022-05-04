package controller

import (
	"net/http"

	"go-api/pkg/utility"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := map[string]string{"message": "OK", "version": "1.0.0"}

		utility.RespondJSON(w, http.StatusOK, body)
	}
}
