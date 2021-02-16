package rio

import (
	"fmt"
)

const (
	// Color Constants

	// Black ...
	Black = "\033[1;30m%s\033[0m"
	// Red ...
	Red = "\033[1;31m%s\033[0m"
	// Green ...
	Green = "\033[1;32m%s\033[0m"
	// Yellow ...
	Yellow = "\033[1;33m%s\033[0m"
	// White ...
	White = "\033[1;37m%s\033[0m"
	// Teal ...
	Teal = "\033[1;36m%s\033[0m"
	// Purple ...
	Purple = "\033[1;34m%s\033[0m"
	// Magenta ...
	Magenta = "\033[1;35m%s\033[0m"
)

// Ignore ...
func Ignore(data interface{}) interface{} {
	return fmt.Sprintf(Black, fmt.Sprintf("%v", data))
}

// Error ...
func Error(data interface{}) interface{} {
	return fmt.Sprintf(Red, fmt.Sprintf("%v", data))
}

// Success ...
func Success(data interface{}) interface{} {
	return fmt.Sprintf(Green, fmt.Sprintf("%v", data))
}

// Warn ...
func Warn(data interface{}) interface{} {
	return fmt.Sprintf(Yellow, fmt.Sprintf("%v", data))
}

// Log ...
func Log(data interface{}) interface{} {
	return fmt.Sprintf(White, fmt.Sprintf("%v", data))
}

// Info ...
func Info(data interface{}) interface{} {
	return fmt.Sprintf(Teal, fmt.Sprintf("%v", data))
}

// Question ...
func Question(data interface{}) interface{} {
	return fmt.Sprintf(Magenta, fmt.Sprintf("%v", data))
}

// Quote ...
func Quote(data interface{}) interface{} {
	return fmt.Sprintf(Purple, fmt.Sprintf("%v", data))
}
