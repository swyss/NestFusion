package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	// Postgres
	PostgresDNS string

	// Influx
	InfluxURL           string
	InfluxToken         string
	InfluxDB            string
	InfluxAdminUser     string
	InfluxAdminPassword string

	// Redis
	RedisURL string
}

func GetEnvironmentVariables(productionFlag *bool) *EnvironmentConfig {
	envFile := ".env.develop"

	if productionFlag != nil && *productionFlag {
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Failed to load .env file")
	}

	return &EnvironmentConfig{
		PostgresDNS:         os.Getenv("DATABASE_URL"),
		InfluxURL:           os.Getenv("INFLUXDB_URL"),
		InfluxToken:         os.Getenv("INFLUXDB_TOKEN"),
		InfluxDB:            os.Getenv("INFLUXDB_DB"),
		InfluxAdminUser:     os.Getenv("INFLUXDB_ADMIN_USER"),
		InfluxAdminPassword: os.Getenv("INFLUXDB_ADMIN_PASSWORD"),
		RedisURL:            os.Getenv("REDIS_URL"),
	}
}
