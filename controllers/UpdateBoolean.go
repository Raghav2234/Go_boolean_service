package controllers

import (
	"Go_boolean_service/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//UpdateBoolean updates the boolean value after Authorisation with token and given valid ID
func UpdateBoolean(database *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		if AuthCheck(c) {
			return
		}
		ID := c.Param("id")
		Value, _ := strconv.ParseBool(c.PostForm("value"))
		db.UpdateBoolean(database, ID, c.PostForm("key"), Value)

	}
}
