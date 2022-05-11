package auth0

import (
	"net/http"

	"go-api/pkg/utility"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

const (
	READABLE  = "read:users"
	WRITEABLE = "write:users"
)

// ユーザーが実行権限（スコープ）を所持しているか検証を行う
// expectedScope: プロセス実行に必要なスコープ
func UseScope(handler http.HandlerFunc, expectedScope string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)

		if claims.HasScope(expectedScope) {
			handler(w, r)
		} else {
			utility.RespondJSON(w, http.StatusUnauthorized, map[string]string{
				"message": "Insufficient scope.",
			})
		}
	}
}
