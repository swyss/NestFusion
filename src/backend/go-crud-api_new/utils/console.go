package utils

import (
	"github.com/fatih/color"
)

// PrintInfo prints a general info message in cyan.
func PrintInfo(format string, a ...interface{}) {
	color.Cyan(format, a...)
}

// PrintSuccess prints a success message in green.
func PrintSuccess(format string, a ...interface{}) {
	color.Green(format, a...)
}

// PrintError prints an error message in red.
func PrintError(format string, a ...interface{}) {
	color.Red(format, a...)
}

// PrintWarning prints a warning message in yellow.
func PrintWarning(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

// FormatInfo returns a formatted info string in cyan.
func FormatInfo(format string, a ...interface{}) string {
	return color.CyanString(format, a...)
}

// FormatSuccess returns a formatted success string in green.
func FormatSuccess(format string, a ...interface{}) string {
	return color.GreenString(format, a...)
}

// FormatError returns a formatted error string in red.
func FormatError(format string, a ...interface{}) string {
	return color.RedString(format, a...)
}

// FormatWarning returns a formatted warning string in yellow.
func FormatWarning(format string, a ...interface{}) string {
	return color.YellowString(format, a...)
}
