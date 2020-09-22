package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateAccessToken generates the access token provided the ID of the boolean
func GenerateAccessToken(ID string) (string, error) {

	expirationTimeAccessToken := time.Now().Add(60 * time.Minute).Unix()
	AtJwtKey := []byte("my_secret_key")

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "signin_1"
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = ID
	claims["exp"] = expirationTimeAccessToken

	// Create the JWT string
	tokenString, err := token.SignedString(AtJwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
