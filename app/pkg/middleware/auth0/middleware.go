package auth0

import (
	"net/http"
	"net/url"
	"time"

	"go-api/pkg/logger"
	"go-api/pkg/utility"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// JWTのバリデーションを行うミドルウェア
func NewMiddleware() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + DOMAIN + "/")
	if err != nil {
		logger.Fatal("Failed to parse the issuer url: " + err.Error())
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{AUDIENCE},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		logger.Fatal("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("Encountered error while validating JWT: " + err.Error())
		utility.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"message": "Encountered error while validating jwt.",
		})
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}
