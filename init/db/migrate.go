package main

import (
	"example.com/m/config"
	"example.com/m/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

func main() {
	db := config.DB

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202206311426",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "202208073215",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Order{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("orders").Error
			},
		},
		{
			ID: "202208073217",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Balance{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("balances").Error
			},
		},
		{
			ID: "202208073210",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Withdrawals{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("withdrawals").Error
			},
		},
	})

	err := m.Migrate()

	config.DB.Model(&domain.Order{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	config.DB.Model(&domain.Balance{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	config.DB.Model(&domain.Withdrawals{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	config.DB.Model(&domain.Withdrawals{}).AddForeignKey("order", "orders(number)", "CASCADE", "CASCADE")

	if err == nil {
		println("Migration did run successfully")
	} else {
		println("Could not migrate: %v", err.Error())
	}
}
