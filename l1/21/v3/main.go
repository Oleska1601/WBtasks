package main

import (
	"fmt"
	"time"
)

// практическая демонстрация на упрощенном реальном примере
// разные логгеры хотим привести к единому интерфейсу

type SimpleLogger struct {
}

func (l *SimpleLogger) LogSimpleMessage(message string) {
	fmt.Println(message, time.Now().Format("15:04:05"))
}

type AdvancedLogger struct {
}

func (l *AdvancedLogger) LogAdvancedMessage(message string) {
	fmt.Println(message, time.Now().Format("2006-01-02 15:04:05"))
}

// целевой интерфейс - Target
type Logger interface {
	Log(message string)
}

type SimpleLoggerAdapter struct {
	*SimpleLogger
}

func NewSimpleLoggerAdapter(logger *SimpleLogger) Logger {
	return &SimpleLoggerAdapter{
		logger,
	}
}

func (a *SimpleLoggerAdapter) Log(message string) {
	a.LogSimpleMessage(message)
}

type AdvancedLoggerAdapter struct {
	*AdvancedLogger
}

func NewAdvancedLoggerAdapter(logger *AdvancedLogger) Logger {
	return &AdvancedLoggerAdapter{
		logger,
	}
}

func (a *AdvancedLoggerAdapter) Log(message string) {
	a.LogAdvancedMessage(message)
}

func main() {
	simpleLogger := &SimpleLogger{}
	advancedLogger := &AdvancedLogger{}
	loggers := []Logger{NewSimpleLoggerAdapter(simpleLogger), NewAdvancedLoggerAdapter(advancedLogger)}
	for _, logger := range loggers {
		logger.Log("some message")
	}
}
