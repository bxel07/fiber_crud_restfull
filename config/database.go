package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB


func ConnectDB()  {
	dsn := "root:Todokana1ko!@tcp(127.0.0.1:3306)/blogs?charset=utf8mb4&parseTime=True&loc=Local"
	
	// disable log for db
	dbLoggers := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      logger.Error,    // Log level (only errors)
			Colorful:      false,           // Disable color
		},
	)
	
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLoggers,
	})

	if err != nil{
		log.Fatal("Failed to connect to database : ", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	
	fmt.Println("Connected to database")
	DB = db
}