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
// Use "Ignore" when you want to print something that is not important,
// or you don't want the text to be highlighted.
func Ignore(data interface{}) interface{} {
	return fmt.Sprintf(Black, fmt.Sprintf("%v", data))
}

// Error ...
// Use "Error" when you want to print some error
// or indicate that something went wrong.
func Error(data interface{}) interface{} {
	return fmt.Sprintf(Red, fmt.Sprintf("%v", data))
}

// Success ...
// Use "Success" when you want to print some success message
// or use it whenever theres some work done and you want to log it.
func Success(data interface{}) interface{} {
	return fmt.Sprintf(Green, fmt.Sprintf("%v", data))
}

// Warn ...
// Use "Warn" when you want to print some issues
// or you want to log something that can raise a problem later.
func Warn(data interface{}) interface{} {
	return fmt.Sprintf(Yellow, fmt.Sprintf("%v", data))
}

// Log ...
// Use "Log" when you want to print some plain text
func Log(data interface{}) interface{} {
	return fmt.Sprintf(White, fmt.Sprintf("%v", data))
}

// Info ...
// Use "Info" when you want to share any information
func Info(data interface{}) interface{} {
	return fmt.Sprintf(Teal, fmt.Sprintf("%v", data))
}

// Question ...
// Use "Question" when you want to raise a question.
func Question(data interface{}) interface{} {
	return fmt.Sprintf(Magenta, fmt.Sprintf("%v", data))
}

// Quote ...
// Use "Quote" when you want print out a quote.
// or you want to share someone else's speech.
func Quote(data interface{}) interface{} {
	return fmt.Sprintf(Purple, fmt.Sprintf("%v", data))
}
