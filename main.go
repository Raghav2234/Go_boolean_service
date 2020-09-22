package main

import (
	"Go_boolean_service/controllers"
	"Go_boolean_service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.CreateConnection()
	if err != nil {
		panic("Database not connected")
	}
	r := gin.Default()
	r.GET("/:id", controllers.GetBoolean(database))
	r.POST("/", controllers.CreateBoolean(database))
	r.PATCH("/:id", controllers.UpdateBoolean(database))
	r.DELETE("/:id", controllers.DeleteBoolean(database))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
