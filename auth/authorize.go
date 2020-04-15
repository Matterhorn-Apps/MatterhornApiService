package auth

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

// CheckScope verifies that an access token that has already been validated grants access to the given scope
func CheckScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})

	claims, ok := token.Claims.(*CustomClaims)

	hasScope := false
	if ok && token.Valid {
		result := strings.Split(claims.Scope, " ")
		for i := range result {
			if result[i] == scope {
				hasScope = true
			}
		}
	}

	return hasScope
}

// GetScopes parses the claims of the given token and returns the scopes string values
func GetScopes(tokenString string) ([]string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			return nil, err
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("unable to validate token claims")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return strings.Split(claims.Scope, " "), nil
}
