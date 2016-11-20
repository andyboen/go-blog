package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=GOBLOG_HOST user=GOBLOG_USER dbname=GOBLOG_DATABASE sslmode=disable password=GOBLOG_PASSWORD")
	if err != nil {
		fmt.Println("Failed to connect to postgres", err)
	} else {
		fmt.Println("Connected to database")
	}
	return db, nil
}
