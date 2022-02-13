package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "sephson:Iama_Winner467/Books?charset=utf&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}

	db = d

} 

func GetDB() *gorm.DB{
	return db
}