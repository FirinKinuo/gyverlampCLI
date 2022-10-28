package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	OutputLogger = log.New(os.Stdout, "", 0)                      // OutputLogger sending messages to stdout
	ErrorLogger  = log.New(os.Stderr, "\033[31mERRO:\033[0m ", 0) // ErrorLogger sending error messages to stderr
	InfoLogger   = log.New(os.Stderr, "\033[32mINFO:\033[0m ", 0) // InfoLogger sending info messages to stderr
)

// Outputf prints a formatted message to stdout
func Outputf(format string, v ...any) {
	OutputLogger.Printf(format, v...)
}

// Output prints a message to stdout
func Output(v ...any) {
	Outputf(fmt.Sprint(v...))
}

// Outputln prints a message with newline to stdout
func Outputln(v ...any) {
	Outputf(fmt.Sprintln(v...))
}

// Errorf prints an error message with formatting to stderr
func Errorf(format string, v ...any) {
	ErrorLogger.Printf(format, v...)
}

// Error prints an error message to stderr
func Error(v ...any) {
	Errorf(fmt.Sprint(v...))
}

// Errorln prints an error message with newline to stderr
func Errorln(v ...any) {
	Errorf(fmt.Sprintln(v...))
}

// Infof prints an info message with formatting to stderr
func Infof(format string, v ...any) {
	InfoLogger.Printf(format, v...)
}

// Info prints an info message to stderr
func Info(v ...any) {
	Infof(fmt.Sprint(v...))
}

// Infoln prints an info message with newline to stderr
func Infoln(v ...any) {
	Infof(fmt.Sprintln(v...))
}
