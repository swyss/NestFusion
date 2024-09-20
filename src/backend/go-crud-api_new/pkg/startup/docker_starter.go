package startup

import (
	"go-crud-api/utils"
	"log"
	"os/exec"
)

// IsDockerRunning checks if Docker is running by executing "docker ps"
func IsDockerRunning() bool {
	cmd := exec.Command("docker", "ps")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// StartDockerEnvironment starts Docker containers if not running
func StartDockerEnvironment() {
	// Verwende FormatInfo anstelle von PrintInfo
	utils.StartSpinner(utils.FormatInfo, "Starting Docker containers")

	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to start Docker containers")
		log.Fatalf("Failed to start Docker containers: %v", err)
	}
	utils.StopSpinner()
	utils.PrintSuccess("Docker environment started successfully.")
}

// ResetDockerEnvironment stops and removes all Docker containers
func ResetDockerEnvironment() {
	// Verwende FormatWarning anstelle von PrintWarning
	utils.StartSpinner(utils.FormatWarning, "Resetting Docker environment")

	cmd := exec.Command("docker-compose", "down", "--volumes", "--remove-orphans")
	if err := cmd.Run(); err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to reset Docker environment")
		log.Fatalf("Failed to reset Docker environment: %v", err)
	}
	utils.StopSpinner()
	utils.PrintWarning("Docker environment reset.")
	StartDockerEnvironment()
}
