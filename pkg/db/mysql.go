package db

import (
	"fmt"

	"main/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"
	host     = "localhost"
	port     = "3306"
	dbname   = "mentorixer"
)


func Init() (*gorm.DB) {
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	err = db.Migrator().DropTable(&model.User{}, &model.Task{}, &model.Friend{}, &model.Room{}, &model.Message{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Task{}, &model.Friend{}, &model.Room{}, &model.Message{})
	if err != nil {
		panic(err)
	}

	return db
}
