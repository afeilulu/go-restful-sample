package database

import (
	"fmt"
	"log"
	"strconv"

	"afeilulu.com/example/config"
	"afeilulu.com/example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Mysql connect string
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Asia%%2fShanghai", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migration to create tables for Order and Item schema
	DB.AutoMigrate(&model.Group{}, &model.User{})
}
