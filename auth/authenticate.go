package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

const auth0ApiIdentifier = "matterhorn-api"
const auth0Domain = "https://matterhorn-prototype.auth0.com/"

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// getPemCerts gets the remote JWKS for the Auth0 account and returns the certificate with the public key in
// PEM format
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(fmt.Sprintf("%s.well-known/jwks.json", auth0Domain))

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

// BuildAuthenticationMiddleware creates a Negroni middleware function that can be used to protect any endpoint
func BuildAuthenticationMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: validationKeyGetter,
		SigningMethod:       jwt.SigningMethodRS256,
	})
}

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	aud := auth0ApiIdentifier
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("invalid audience")
	}
	// Verify 'iss' claim
	iss := auth0Domain
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("invalid issuer")
	}

	cert, err := getPemCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}
