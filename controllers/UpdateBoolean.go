package controllers

import (
	"Go_boolean_service/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//UpdateBoolean updates the boolean value after Authorisation with token and given valid ID
func UpdateBoolean(database *gorm.DB) func(*gin.Context) {
	var boolObj db.BooleanTemp
	return func(c *gin.Context) {
		if AuthCheck(c) {
			return
		}
		ID := c.Param("id")
		if err := c.ShouldBindJSON(&boolObj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.UpdateBoolean(database, ID, boolObj)
		boolObj := db.ReadBoolean(database, ID)
		c.JSON(http.StatusOK, gin.H{"id": boolObj.Id, "key": boolObj.Key, "value": boolObj.Value})
	}
}
