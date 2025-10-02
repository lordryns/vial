package logger

import (
	"fmt"
	"unicode/utf8"

	"github.com/fatih/color"
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
		"[%v] - %v\n",
		decideColor(m_type),
		message,
	)
}

// this is redundant i know
func decideColor(m_type string) string {
	var cString string

	switch m_type {
	case "error":
		cString = color.RedString(m_type)
	case "warning":
		cString = color.YellowString(m_type)
	case "info":
		cString = color.BlueString(m_type)
	case "success":
		cString = color.GreenString(m_type)
	}

	return cString
}
