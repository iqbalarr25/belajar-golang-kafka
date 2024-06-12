package config

import (
	"os"
	"strconv"
)

type App struct {
	Name      string
	Env       string
	Host      string
	ImageName string
}

func GetApp() App {
	return App{
		Name:      GetEnv("APP_NAME", "Golang"),
		Env:       GetEnv("APP_ENV", "local"),
		Host:      GetEnv("APP_HOST", "localhost"),
		ImageName: GetEnv("IMAGE_NAME", ""),
	}
}

type Cache struct {
	Host     string
	Port     string
	Database int
	Password string
	Lifetime string
}

func GetCache() Cache {
	database, _ := strconv.Atoi(GetEnv("REDIS_DATABASE", "0"))
	return Cache{
		Host:     GetEnv("REDIS_HOST", "127.0.0.1"),
		Port:     GetEnv("REDIS_PORT", "6379"),
		Database: database,
		Password: GetEnv("REDIS_PASSWORD", ""),
		Lifetime: GetEnv("REDIS_LIFETIME", "30s"),
	}
}

type Database struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Timezone string
}

func GetDatabase() Database {
	return Database{
		Host:     GetEnv("DB_HOST", "127.0.0.1"),
		Port:     GetEnv("DB_PORT", "5432"),
		Database: GetEnv("DB_DATABASE", "golang"),
		Username: GetEnv("DB_USERNAME", "root"),
		Password: GetEnv("DB_PASSWORD", ""),
		Timezone: GetEnv("TIMEZONE", "Asia/Jakarta"),
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
