package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is database instance
var DB *gorm.DB

func init() {
	connectionString, err := gorm.Open(Env.BdConnection)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("You connected to your database.")
}
