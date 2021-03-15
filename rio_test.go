package rio

import (
	"fmt"
	"testing"
)

const (
	black   = "\033[1;30m%s\033[0m"
	red     = "\033[1;31m%s\033[0m"
	green   = "\033[1;32m%s\033[0m"
	yellow  = "\033[1;33m%s\033[0m"
	white   = "\033[1;37m%s\033[0m"
	teal    = "\033[1;36m%s\033[0m"
	purple  = "\033[1;34m%s\033[0m"
	magenta = "\033[1;35m%s\033[0m"
)

func TestIgnore(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Ignore("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Ignore function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Ignore("Ignore")) != len(fmt.Sprintf(black, fmt.Sprintf("%v", "Ignore"))) {
		t.Fail()
	}
}

func TestError(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Error("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Error function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Error("Error")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Error"))) {
		t.Fail()
	}
}

func TestSuccess(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Success("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Success function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Success("Success")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Success"))) {
		t.Fail()
	}
}

func TestWarn(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Warn("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Warn function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Warn("Warn")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Warn"))) {
		t.Fail()
	}
}

func TestLog(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Log("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Log function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Log("Log")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Log"))) {
		t.Fail()
	}
}

func TestInfo(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Info("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Info function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Info("Info")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Info"))) {
		t.Fail()
	}
}

func TestQuestion(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Question("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Question function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Question("Question")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Question"))) {
		t.Fail()
	}
}

func TestQuote(t *testing.T) {
	// ANSII code counts, the length returns 11, even though the string is empty
	// That's how we know theres some hidden codes that change the color of the text.
	if !(len(Quote("")) > 0) {
		t.Fail()
	}
	// Here we are calling the Quote function and checking the length with
	// manually experimenting with fmt.Sprintf
	if len(Quote("Quote")) != len(fmt.Sprintf(red, fmt.Sprintf("%v", "Quote"))) {
		t.Fail()
	}
}
