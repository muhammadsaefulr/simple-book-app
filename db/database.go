package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadsaefulr/simple-book-app/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error while loading the env file")
	}

	DBusername := os.Getenv("DB_USERNAME")
	DBpassword := os.Getenv("DB_PASSWORD")
	DBhost := os.Getenv("DB_HOST")
	DBport := os.Getenv("DB_PORT")
	DBname := os.Getenv("DB_NAME")

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBusername, DBpassword, DBhost, DBport, DBname)
	db, err := gorm.Open(mysql.Open(dburl), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(models.Detail{})

	return db, nil
}
