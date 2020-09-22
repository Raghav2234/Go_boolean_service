package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthCheck checks the validity of token
func AuthCheck(c *gin.Context) bool {
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		fmt.Printf("Authorization token was not provided")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
		c.Abort()
		return true
	}
	claims := jwt.MapClaims{}
	AtJwtKey := []byte("my_secret_key")
	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		fmt.Printf("Incorrect Format of Authn Token")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Format of Authorization Token "})
		c.Abort()
		return true
	}
	parsedToken, err := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
		return AtJwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Printf("Invalid Token Signature")
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token Signature"})
			c.Abort()
			return true
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad Request"})
		c.Abort()
		return true
	}
	if !parsedToken.Valid {
		fmt.Printf("Invalid Token")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
		c.Abort()
		return true
	}

	return false
}
