package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() *gorm.DB {

	// Get database details from environment details
	host := GetDatabaseHost()
	port := GetDatabasePort()
	user := GetDatabaseUser()
	name := GetDatabaseUser()
	password := GetDatabasePassword()

	connectionString := fmt.Sprintf(
		"postgresql: //$%s:$%s@$%s:$%s/$%s?sslmode=disable",
		user, password, host, port, name,
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
