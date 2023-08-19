package applogger

import "fmt"

type consoleLogger struct {
	name string
}

func (logger consoleLogger) I(message any, args ...any) {
	fmt.Printf("[INFO] [%s] %v\n", logger.name, message)
}

func (logger consoleLogger) E(message any, args ...any) {
	fmt.Printf("[ERROR] [%s] %v\n", logger.name, message)
}

func (logger consoleLogger) D(message any, args ...any) {
	fmt.Printf("[DEBUG] [%s] %v\n", logger.name, message)
}

func newConsoleLogger(name string) consoleLogger {
	return consoleLogger{
		name: name,
	}
}
