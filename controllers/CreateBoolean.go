package controllers

import (
	"Go_boolean_service/auth"
	"Go_boolean_service/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateBoolean creates the boolean provided key-value pair and returns the identifier(ID)
func CreateBoolean(database *gorm.DB, jobChan chan db.Boolean) func(*gin.Context) {
	return func(c *gin.Context) {
		var obj db.BooleanTemp
		if err := c.ShouldBindJSON(&obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ID := db.CreateUUID()
		boolObj := db.Boolean{Id: ID, Key: obj.Key, Value: obj.Value}
		jobChan <- boolObj
		// resp := db.CreateBoolean(database, boolObj)
		token, _ := auth.GenerateAccessToken(ID)
		// obj := db.ReadBoolean(database, resp.Id)
		// fmt.Println(obj.Id, obj.Key, obj.Value)
		c.JSON(http.StatusOK, gin.H{"token": token, "id": ID, "key": boolObj.Key, "value": boolObj.Value})
	}
}
