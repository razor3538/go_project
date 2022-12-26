package config

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is database instance
var DB *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open(Env.BdConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("You connected to your database.")
}
