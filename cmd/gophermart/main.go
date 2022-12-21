package main

import (
	_ "example.com/m/docs"
	"example.com/m/internal/routes"
)

// @title    Gophermart Service
// @version  1.0.0
// @host     localhost:8000
func main() {
	r := routes.SetupRouter()

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
