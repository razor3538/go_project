package main

import (
	"example.com/m/config"
	_ "example.com/m/docs"
	"example.com/m/internal/routes"
)

// @title    Gophermart Service
// @version  1.0.0
// @host     localhost:8000
func main() {
	config.CheckFlagEnv()
	config.InitDB()

	r := routes.SetupRouter()

	if err := r.Run(config.Env.Address); err != nil {
		panic(err)
	}
}
