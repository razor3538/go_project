package config

import (
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Port       string
	Host       string
}

// Env is env project struct
var Env env

func init() {
	_ = godotenv.Load(".env.example")
	Env = env{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
	}
}
