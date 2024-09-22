package startup

import (
	"go-crud-api/utils"
	"log"
	"os/exec"
	"strings"
	"time"
)

// IsDockerRunning checks if Docker is running by executing "docker ps"
func IsDockerRunning() bool {
	cmd := exec.Command("docker", "ps")
	if err := cmd.Run(); err != nil {
		utils.PrintWarning("Docker is not running.")
		return false
	}
	return true
}

// BuildDockerContainers builds the Docker containers if necessary
func BuildDockerContainers() {
	utils.StartSpinner(utils.FormatInfo, "Building Docker containers")

	// Execute the command to build Docker containers
	cmd := exec.Command("docker-compose", "up", "--build", "-d")
	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to build Docker containers. Output: " + string(output))
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintSuccess("Docker containers built and started successfully.")
}

// StartDockerEnvironment starts Docker containers if not running
func StartDockerEnvironment() {
	utils.StartSpinner(utils.FormatInfo, "Starting Docker containers")

	// Start the containers using docker-compose
	cmd := exec.Command("docker-compose", "up", "-d")
	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to start Docker containers. Output: " + string(output))
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintSuccess("Docker environment started successfully.")

	// Display the list of running containers
	ListRunningContainers()

	// Optional: Wait for containers to be fully initialized
	time.Sleep(5 * time.Second)
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

// RemoveDockerImages removes all unused Docker images
func RemoveDockerImages() {
	utils.StartSpinner(utils.FormatWarning, "Removing unused Docker images")

	// Remove all stopped containers and unused images
	cmd := exec.Command("docker", "system", "prune", "-af")
	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to remove Docker images. Output: " + string(output))
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintSuccess("Unused Docker images removed.")
}

// ResetDockerEnvironment stops and removes all Docker containers and images
func ResetDockerEnvironment() {
	utils.StartSpinner(utils.FormatWarning, "Resetting Docker environment")

	// Stop and remove all containers, volumes, and orphaned containers
	cmd := exec.Command("docker-compose", "down", "--volumes", "--remove-orphans")
	output, err := cmd.CombinedOutput() // Capture both stdout and stderr
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to reset Docker environment. Output: " + string(output))
		log.Fatalf("Error: %v", err)
	}
	utils.StopSpinner()
	utils.PrintWarning("Docker environment reset.")

	// Remove all unused Docker images
	RemoveDockerImages()

	// Build and start the Docker environment again
	BuildDockerContainers()
}
