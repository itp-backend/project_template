package config

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName      string
	AppPort      int
	LogLevel     string
	Environment  string
	JWTSecret    string
	RedisAddress string
	DBUsername   string
	DBPassword   string
	DBHost       string
	DBPort       int
	DBName       string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Warning("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:      GetString("APP_NAME"),
		AppPort:      GetInt("APP_PORT"),
		LogLevel:     GetString("LOG_LEVEL"),
		Environment:  GetString("ENVIRONMENT"),
		JWTSecret:    GetString("JWT_SECRET"),
		RedisAddress: GetString("REDIS_ADDRESS"),
		DBUsername:   GetString("DB_USERNAME"),
		DBPassword:   GetString("DB_PASSWORD"),
		DBHost:       GetString("DB_HOST"),
		DBPort:       GetInt("DB_PORT"),
		DBName:       GetString("DB_NAME"),
	}

	return appConfig
}
