package utility

import "github.com/sirupsen/logrus"

//LoggerInterface ...
type LoggerInterface interface {
	Info(string)
	Error(string)
}

//NewLogrusLogger ...
func NewLogrusLogger() *Logger {
	return &Logger{logger: logrus.New()}
}

//Logger is implements LoggerInterface
type Logger struct {
	logger *logrus.Logger
}

//Info logger ...
func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

//Error logger ...
func (l *Logger) Error(message string) {
	l.logger.Error(message)
}

//LoggerStub is stub for LoggerInterface
type LoggerStub struct {
}

//Info loggerStub ...
func (l *LoggerStub) Info(message string) {
}

//Error loggerStub
func (l *LoggerStub) Error(message string) {
}
