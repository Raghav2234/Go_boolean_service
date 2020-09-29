package db

import (
	"Go_boolean_service/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//CreateConnection creates the database connection
func CreateConnection() (*gorm.DB, error) {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var user, password, dbName, dbHost string
	var dbPort int
	// if _, err := os.Stat("/.dockerenv"); err == nil {
	// user = os.Getenv("DOCK_USR")
	// password = os.Getenv("DOCK_PASS")
	// dbName = os.Getenv("DOCK_DB")
	// dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	// } else {
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")
	dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	// }

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	fmt.Println("Database successfully connected")
	db.AutoMigrate(&models.Boolean{})

	return db, nil
}
