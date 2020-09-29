package queue

import (
	"Go_boolean_service/models"

	"gorm.io/gorm"
)

//MessageQueue stores inMemory queue to make creation flow aync
func MessageQueue(database *gorm.DB) chan models.Boolean {
	jobChan := make(chan models.Boolean, 100)
	for i := 0; i < 5; i++ {
		go Worker(jobChan, database)
	}
	return jobChan
}
