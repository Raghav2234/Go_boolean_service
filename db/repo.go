package db

import (
	"Go_boolean_service/models"
	"crypto/rand"
	"fmt"
	"log"

	"gorm.io/gorm"
)

//CreateUUID creates UUID
func CreateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

//CreateBoolean creates the boolean object
func CreateBoolean(db *gorm.DB, boolObj models.Boolean) models.Boolean {
	// Create
	// ID := createUUID()
	db.Create(&boolObj)
	return boolObj
}

//ReadBoolean reads the boolean from the database using Id field
func ReadBoolean(db *gorm.DB, ID string) models.Boolean {
	var boolObj models.Boolean
	// db.First(&boolObj, ID)
	db.Where("id = ?", ID).First(&boolObj)
	return boolObj
}

//UpdateBoolean updates the boolean parameters
func UpdateBoolean(db *gorm.DB, ID string, obj models.BooleanTemp) {
	var boolObj models.Boolean
	db.Where("id = ?", ID).First(&boolObj)
	db.Model(&boolObj).Updates(models.Boolean{Key: obj.Key, Value: obj.Value})

}

//DeleteBoolean deletes the boolean entry identified by Id
func DeleteBoolean(db *gorm.DB, ID string) {
	// db.Delete(&Boolean{}, ID)
	db.Where("id = ?", ID).Delete(&models.Boolean{})

}
