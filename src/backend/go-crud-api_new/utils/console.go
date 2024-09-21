package utils

import (
	"github.com/fatih/color"
)

// PrintInfo prints a general info message in bold cyan.
func PrintInfo(format string, a ...interface{}) {
	color.New(color.FgCyan, color.Bold).Printf(format+"\n", a...)
}

// PrintSuccess prints a success message in bold green.
func PrintSuccess(format string, a ...interface{}) {
	color.New(color.FgGreen, color.Bold).Printf(format+"\n", a...)
}

// PrintError prints an error message in bold red.
func PrintError(format string, a ...interface{}) {
	color.New(color.FgRed, color.Bold).Printf(format+"\n", a...)
}

// PrintWarning prints a warning message in bold yellow.
func PrintWarning(format string, a ...interface{}) {
	color.New(color.FgYellow, color.Bold).Printf(format+"\n", a...)
}

// FormatInfo returns a formatted info string in bold cyan.
func FormatInfo(format string, a ...interface{}) string {
	return color.New(color.FgCyan, color.Bold).Sprintf(format, a...)
}

// FormatSuccess returns a formatted success string in bold green.
func FormatSuccess(format string, a ...interface{}) string {
	return color.New(color.FgGreen, color.Bold).Sprintf(format, a...)
}

// FormatError returns a formatted error string in bold red.
func FormatError(format string, a ...interface{}) string {
	return color.New(color.FgRed, color.Bold).Sprintf(format, a...)
}

// FormatWarning returns a formatted warning string in bold yellow.
func FormatWarning(format string, a ...interface{}) string {
	return color.New(color.FgYellow, color.Bold).Sprintf(format, a...)
}
