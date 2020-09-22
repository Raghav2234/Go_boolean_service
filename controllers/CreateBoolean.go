package controllers

import (
	"Go_boolean_service/auth"
	"Go_boolean_service/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateBoolean creates the boolean provided key-value pair and returns the identifier(ID)
func CreateBoolean(database *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		if (strings.ToLower(c.PostForm("value")) != "true") && (strings.ToLower(c.PostForm("value")) != "false") {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "bool value required"})
			return
		}
		value, _ := strconv.ParseBool(strings.ToLower(c.PostForm("value")))
		boolObj := db.CreateBoolean(database, c.PostForm("key"), value)
		token, _ := auth.GenerateAccessToken(boolObj.Id)
		c.JSON(http.StatusOK, gin.H{"token": token, "id": boolObj.Id, "key": boolObj.Key, "value": boolObj.Value})
	}
}
