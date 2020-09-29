package main

import (
	"Go_boolean_service/controllers"
	"Go_boolean_service/db"
	"Go_boolean_service/queue"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	database, err := db.CreateConnection()
	if err != nil {
		panic("Database not connected")
	}
	serverSetup(database).Run()

}
func serverSetup(database *gorm.DB) *gin.Engine {

	jobChan := queue.MessageQueue(database)

	r := gin.Default()
	r.GET("/:id", controllers.GetBoolean(database))
	r.POST("/", controllers.CreateBoolean(database, jobChan))
	r.PATCH("/:id", controllers.UpdateBoolean(database))
	r.DELETE("/:id", controllers.DeleteBoolean(database))
	return r // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
