package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go-crud-api/utils"
	"log"
	"os"
)

// InitializeInfluxDB initializes the InfluxDB client
func InitializeInfluxDB() influxdb2.Client {
	influxURL := os.Getenv("INFLUXDB_URL")
	if influxURL == "" {
		utils.PrintError("INFLUXDB_URL environment variable not set")
		log.Fatal("INFLUXDB_URL environment variable not set")
	}
	utils.PrintSuccess("INFLUXDB_URL loaded successfully")

	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		utils.PrintError("INFLUXDB_TOKEN environment variable not set")
		log.Fatal("INFLUXDB_TOKEN environment variable not set")
	}
	utils.PrintSuccess("INFLUXDB_TOKEN loaded successfully")

	InfluxClient := influxdb2.NewClient(influxURL, token)

	utils.StartSpinner(utils.FormatInfo, "Connecting to InfluxDB")

	// Ping InfluxDB
	_, err := InfluxClient.Ready(context.Background())
	if err != nil {
		utils.StopSpinnerWithError("InfluxDB")
		log.Fatalf("Failed to connect to InfluxDB: %v", err)
	} else {
		utils.StopSpinner()
		utils.PrintSuccess("Successfully connected to InfluxDB")
	}

	return InfluxClient
}
