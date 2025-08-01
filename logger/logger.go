package logger

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Logger struct {
	LogCount int
}

func NewLogger() *Logger {
	return &Logger{LogCount: 0}
}

func (logger *Logger) Success(message string) {
	baseLog("success", message)
}

func (logger *Logger) Info(message string) {
	baseLog("info", message)
}

func (logger *Logger) Warning(message string) {
	baseLog("warning", message)
}

func (logger *Logger) Error(message string) {
	baseLog("error", message)
}

func baseLog(m_type string, message string) {
	var spaceCount int = utf8.RuneCountInString(m_type)
	spaceCount = spaceCount / 2
	fmt.Printf(
		"[ %v %v] - %v\n",
		strings.ToUpper(m_type), strings.Repeat(" ", spaceCount),
		message,
	)
}
