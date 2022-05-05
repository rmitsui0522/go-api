package auth0

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKS struct {
	Keys []JSONWebKeys `json:"keys"`
}

func getJWKS(auth0Domain string) (*JWKS, error) {
	url := fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	jwks := &JWKS{}
	err = json.NewDecoder(res.Body).Decode(jwks)

	return jwks, err
}

// JWKSに含まれる公開鍵をPEM形式で返す
func createPemCert(jwks *JWKS, token *jwt.Token) (string, error) {
	cert := ""

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return "", errors.New("unable to find appropriate key")
	}

	return cert, nil
}
