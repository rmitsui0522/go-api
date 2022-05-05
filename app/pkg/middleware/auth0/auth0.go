package auth0

import (
	"errors"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func newMiddleware(jwks *JWKS) (*jwtmiddleware.JWTMiddleware, error) {
	options := jwtmiddleware.Options{
		ValidationKeyGetter: newValidationKeyGetter(domain, clientID, jwks),
		SigningMethod:       jwt.SigningMethodRS256,
		ErrorHandler:        func(w http.ResponseWriter, r *http.Request, err string) {},
	}
	return jwtmiddleware.New(options), nil
}

func newValidationKeyGetter(domain, clientID string, jwks *JWKS) func(*jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return token, errors.New("invalid claims type")
		}

		// azpフィールドを見て、適切なClientIDのJWTかチェックする
		azp, ok := claims["azp"].(string)

		if !ok {
			return nil, errors.New("authorized parties are required")
		}
		if azp != clientID {
			return nil, errors.New("invalid authorized parties")
		}

		// issフィールドを見て、正しいトークン発行者か確認する
		iss := fmt.Sprintf("https://%s/", domain)
		ok = token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
		if !ok {
			return nil, errors.New("invalid issuer")
		}

		// JWTの検証に必要な鍵を生成する
		cert, err := createPemCert(jwks, token)
		if err != nil {
			return nil, err
		}

		return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	}
}
