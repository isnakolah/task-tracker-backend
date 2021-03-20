package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() *gorm.DB {

	connectionString := fmt.Sprintf(
		"postgresql: //$%s:$%s@$%s:$%s/$%s?sslmode=disable",
		DbUser, DbPassword, DbHost, DbPort, DbName,
	)

	db, err := gorm.Open("postgres", connectionString)

	if err == nil {
		return db
	}
	panic(err.Error())
}

func GetDB() *gorm.DB {
	return DB
}
