package controllers

import (
	"Go_boolean_service/auth"
	"Go_boolean_service/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateBoolean creates the boolean provided key-value pair and returns the identifier(ID)
func CreateBoolean(database *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var boolObj db.BooleanTemp
		if err := c.ShouldBindJSON(&boolObj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := db.CreateBoolean(database, boolObj)
		token, _ := auth.GenerateAccessToken(resp.Id)
		obj := db.ReadBoolean(database, resp.Id)
		fmt.Println(obj.Id, obj.Key, obj.Value)
		c.JSON(http.StatusOK, gin.H{"token": token, "id": resp.Id, "key": resp.Key, "value": resp.Value})
	}
}
