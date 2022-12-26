package config

import (
	migrate "example.com/m/init/db"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is database instance
var DB *gorm.DB

func init() {
	db, err := gorm.Open("postgres", Env.Address)
	if err != nil {
		panic(err)
	}

	DB = db

	migrate.Migrate(db)

	fmt.Println("You connected to your database.")
}
