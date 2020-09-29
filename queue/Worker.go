package queue

import (
	"Go_boolean_service/db"
	"Go_boolean_service/models"
	"fmt"

	"gorm.io/gorm"
)

// Worker creates queue to make creation flow async
func Worker(jobChan <-chan models.Boolean, database *gorm.DB) {
	for job := range jobChan {
		fmt.Println(job.Id)
		db.CreateBoolean(database, job)
	}
}
