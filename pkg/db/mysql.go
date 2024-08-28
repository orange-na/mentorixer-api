package db

import (
	"fmt"

	"main/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"
	password = "Yuto8181nmb"
	host     = "localhost"
	port     = "3306"
	dbname   = "todos"
)


func Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&model.User{}, &model.Task{})
	db.AutoMigrate(&model.User{}, &model.Task{})
	
	return db, err
}
