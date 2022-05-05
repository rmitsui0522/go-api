package auth0

import (
	"context"
	"net/http"

	"go-api/pkg/utility"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog/log"
)

// jwtmiddleware.JWTMiddlewareをContextに格納するキー
type JWTMiddlewareKey struct {
	*jwtmiddleware.JWTMiddleware
}

// JWTをContextに保存するキー
type JWTKey struct {
	*jwt.Token
}

func NewMiddleware() (func(http.Handler) http.Handler, error) {
	jwks, err := getJWKS(domain)
	if err != nil {
		log.Fatal().Err(err).Msg("Fatal: fetch JSON web keys")
	}

	jwtMiddleware, err := newMiddleware(jwks)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ContextにJWTMiddlewareを格納
			ctx := context.WithValue(r.Context(), JWTMiddlewareKey{}, jwtMiddleware)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}, err
}

// JWT検証を行うためのmiddleware
func UseJWT(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ContextからJWTMiddlewareを取得
		jwtm := r.Context().Value(JWTMiddlewareKey{}).(*jwtmiddleware.JWTMiddleware)

		// リクエスト中のJWTを検証
		if err := jwtm.CheckJWT(w, r); err != nil {
			utility.RespondJSON(w, http.StatusUnauthorized, map[string]string{
				"message": err.Error(),
			})
			return
		}

		// JWT検証後に、Contextのjwtm.Options.UserPropertyからパース済みのトークンを取得する
		if val := r.Context().Value(jwtm.Options.UserProperty); val != nil {
			token, ok := val.(*jwt.Token)
			if ok {
				// リクエストのContextにJWTを保存する
				ctx := context.WithValue(r.Context(), JWTKey{}, token)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func GetJWT(ctx context.Context) *jwt.Token {
	rawJWT, ok := ctx.Value(JWTKey{}).(*jwt.Token)
	if !ok {
		return nil
	}
	return rawJWT
}
