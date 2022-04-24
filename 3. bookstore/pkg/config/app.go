package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Create a DB object instance
var (
	db *gorm.DB
)

//Connect it with the locally run mySQL server using your mySQL username and password
func Connect() {
	d, err := gorm.Open("mysql", "abheisenberg:Nhksmksym@96@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

//Return the DB
func GetDB() *gorm.DB {
	return db
}
