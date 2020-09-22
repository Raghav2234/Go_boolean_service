package controllers

import (
	"Go_boolean_service/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//GetBoolean fetch the boolean from the database with provided ID
func GetBoolean(database *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		ID := c.Param("id")
		boolObj := db.ReadBoolean(database, ID)
		if boolObj.Id == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found."})
			return
		}
		c.JSON(http.StatusFound, gin.H{"id": boolObj.Id, "key": boolObj.Key, "value": boolObj.Value})
	}
}
