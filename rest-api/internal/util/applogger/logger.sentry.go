package applogger

type sentryLogger struct {
	name          string
	consoleLogger consoleLogger
}

func (logger sentryLogger) I(message any, args ...any) {}

func (logger sentryLogger) E(message any, args ...any) {
	logger.consoleLogger.E(message, args)
}

func (logger sentryLogger) D(message any, args ...any) {}

func newSentryLogger(name string) sentryLogger {
	var consoleLogger = newConsoleLogger(name)
	return sentryLogger{
		name:          name,
		consoleLogger: consoleLogger,
	}
}
