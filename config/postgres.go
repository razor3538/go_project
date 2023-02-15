package config

import (
	migrate "example.com/m/init/db"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

// DB is database instance
var DB *gorm.DB

func InitDB() {
	//connectionString := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=tander password=123 sslmode=disable")

	// test
	db, err := gorm.Open("postgres", Env.BdConnection)
	//db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		os.Exit(4)
	}

	DB = db

	migrate.Migrate(db)

	fmt.Println("You connected to your database.")
}
