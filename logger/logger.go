package logger

import (
	"strings"

	"github.com/fatih/color"
)

//Logger defines the shell logger structure
type Logger struct {
	Depth int
}

//DepthIn adds a log level
func (logger *Logger) DepthIn() {
	logger.Depth++
}

//DepthOut adds a log level
func (logger *Logger) DepthOut() {
	if logger.Depth > 0 {
		logger.Depth--
	}
}

//INFO prints an info log to shell
func (logger *Logger) INFO(message string) {
	color.Yellow("%s[INFO]:%s", logger.getIndent(), message)
}

//ERROR prints an error log to shell
func (logger *Logger) ERROR(message string) {
	color.Red("%s[ERROR]:%s", logger.getIndent(), message)
}

//SUCCESS prints an error log to shell
func (logger *Logger) SUCCESS(message string) {
	color.Green("%s[SUCCESS]:%s", logger.getIndent(), message)
}

//getIndent returns an indent string as log prefix
func (logger *Logger) getIndent() string {
	return strings.Repeat("=", logger.Depth)
}
