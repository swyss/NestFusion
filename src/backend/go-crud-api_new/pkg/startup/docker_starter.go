package startup

import (
	"go-crud-api/utils"
	"log"
	"os/exec"
	"strings"
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
	utils.StartSpinner(utils.FormatInfo, "Starting Docker containers")

	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to start Docker containers. Please verify your Docker installation and configuration.")
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintSuccess("Docker environment started successfully.")

	// Display the list of running containers
	ListRunningContainers()
}

// ListRunningContainers fetches and prints the list of running Docker containers
func ListRunningContainers() {
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}: {{.Status}}")
	output, err := cmd.Output()
	if err != nil {
		utils.PrintError("Failed to fetch running Docker containers")
		log.Fatalf("Error: %v", err)
	}

	containerList := strings.TrimSpace(string(output))
	if containerList == "" {
		utils.PrintWarning("No containers are currently running.")
	} else {
		utils.PrintInfo("The following containers are running:")
		utils.PrintSuccess(containerList)
	}
}

// ResetDockerEnvironment stops and removes all Docker containers
func ResetDockerEnvironment() {
	utils.StartSpinner(utils.FormatWarning, "Resetting Docker environment")

	cmd := exec.Command("docker-compose", "down", "--volumes", "--remove-orphans")
	if err := cmd.Run(); err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to reset Docker environment. Ensure Docker is installed and running.")
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintWarning("Docker environment reset.")
	StartDockerEnvironment()
}
