package config

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func BuildDBConfig() (dbconfig string) {

	dbConfig := "root:Trabzonspor.61@tcp(localhost:3306)/tododb?charset=utf8&parseTime=True&loc=Local"

	return dbConfig
}
