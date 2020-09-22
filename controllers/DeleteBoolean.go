package controllers

import (
	"Go_boolean_service/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//DeleteBoolean deletes the boolean provided token and ID
func DeleteBoolean(database *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		if AuthCheck(c) {
			return
		}
		ID := c.Param("id")
		db.DeleteBoolean(database, ID)
		c.JSON(http.StatusNoContent, gin.H{"value": "No content found"})
	}
}
