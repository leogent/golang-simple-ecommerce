package repository

import (
	"log"
	"os"
	"simple-ecommerce/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+"/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db

}
