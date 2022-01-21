package repository

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Error connecting to database")
		return nil
	}

	//db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db

}
