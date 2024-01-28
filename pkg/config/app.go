package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "gnuvse:927339@localhost/db1?icharset=utf8&parseTime=True*loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *form.DB {
	return db
}
