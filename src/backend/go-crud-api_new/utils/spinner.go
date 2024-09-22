package utils

import (
	"fmt"
	"time"
)

var currentSpinnerStopChan chan bool

// StartSpinner starts a new spinner with the given color function and message.
func StartSpinner(colorFn func(format string, a ...interface{}) string, message string) {
	// Stop any running spinner before starting a new one
	StopSpinner()

	currentSpinnerStopChan = make(chan bool)
	go spinner(currentSpinnerStopChan, colorFn, message)
}

// StopSpinner stops the currently running spinner.
func StopSpinner() {
	if currentSpinnerStopChan != nil {
		currentSpinnerStopChan <- true
		currentSpinnerStopChan = nil
	}
}

// spinner logic that runs until the stop signal is sent
func spinner(stopChan chan bool, colorFn func(format string, a ...interface{}) string, message string) {
	chars := `-\|/`
	i := 0

	for {
		select {
		case <-stopChan:
			// Stop the spinner and print the completed message
			fmt.Printf("\r✔ %s completed!\n", colorFn(message))
			return
		default:
			// Show the spinner animation
			fmt.Printf("\r%s %s %c", colorFn(message), colorFn("in progress"), chars[i%len(chars)])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// StopSpinnerWithError stops the spinner and shows an error message.
func StopSpinnerWithError(message string) {
	StopSpinner()
	PrintError(fmt.Sprintf("✘ %s failed!", message))
}
