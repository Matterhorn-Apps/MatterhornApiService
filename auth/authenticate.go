package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

const auth0ApiIdentifier = "matterhorn-api"
const auth0Domain = "https://matterhorn-prototype.auth0.com/"

// getPemCerts gets the remote JWKS for the Auth0 account and returns the certificate with the public key in
// PEM format
func getPemCert(token *jwt.Token) (string, error) {
	type JSONWebKeys struct {
		Kty string   `json:"kty"`
		Kid string   `json:"kid"`
		Use string   `json:"use"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5c []string `json:"x5c"`
	}

	type Jwks struct {
		Keys []JSONWebKeys `json:"keys"`
	}

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

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

// BuildAuthenticationMiddleware creates a middleware function that can be used to protect any HTTP endpoint
func BuildAuthenticationMiddleware() *jwtmiddleware.JWTMiddleware {
	//
	// THIS CODE BYPASSES AUTHENTICATION - DO NOT ENABLE IN PRODUCTION ENVIRONMENT
	//
	// Require valid JWT token to access endpoint only in production environment.
	// This enables us to use GraphQL playground and use fake test accounts without
	// authenticating with Auth0 in local and dev environments.
	credentialsOptional := os.Getenv("MATTERHORN_ENV") != "prod"

	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: validationKeyGetter,
		SigningMethod:       jwt.SigningMethodRS256,
		// We don't currently require authentication for all requests
		// Authorization requirements must be explicitly enforced by the resolvers where necessary
		CredentialsOptional: credentialsOptional,
	})
}

// GetUserIDFromContext retrieves the User ID from a given JWT access token
func GetUserIDFromContext(ctx context.Context) (*string, error) {
	//
	// THIS CODE BYPASSES AUTHENTICATION - DO NOT ENABLE IN PRODUCTION ENVIRONMENT
	//
	// Optionally override User ID in local environment only
	if os.Getenv("MATTERHORN_ENV") != "prod" {
		if testUser, exists := os.LookupEnv("TEST_USER_ID"); exists {
			if !strings.HasPrefix(testUser, "test|") {
				return nil, errors.New("test user id must begin with 'test|'")
			}

			log.Printf("authenticating as test user: %s", testUser)
			return &testUser, nil
		}
	}

	userToken, ok := ctx.Value("user").(*jwt.Token)
	if userToken == nil || !ok {
		return nil, errors.New("request does not contain user token")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if claims == nil || !ok {
		return nil, errors.New("unable to parse claims on user token")
	}

	// 'sub' contains unique user ID assigned by Auth0
	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("unable to read 'sub' value from user token")
	}

	return &sub, nil
}

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	// TODO: Code adapted from Auth0 samples fails to detect audience from valid token #36
	// https://github.com/Matterhorn-Apps/MatterhornApiService/issues/36

	// Verify 'aud' claim
	// Audience is expected to match value for MatterhornAPIService
	aud := auth0ApiIdentifier
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		log.Printf("invalid audience on token: %v", token)
		return token, errors.New("invalid audience")
	}
	// Verify 'iss' claim
	// Issuer is expected to match value for Matterhorn Auth0 tenant
	iss := auth0Domain
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, true)
	if !checkIss {
		log.Printf("invalid issuer on token: %v", token)
		return token, errors.New("invalid issuer")
	}

	cert, err := getPemCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}
