package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents a logging level
type LogLevel int

// Log levels
const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// Logger is a simple logger that writes to multiple outputs
type Logger struct {
	logLevel  LogLevel
	outputs   []io.Writer
	showDate  bool
	showFile  bool
	dateFormat string
}

// NewLogger creates a new logger
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		logLevel:   level,
		outputs:    []io.Writer{os.Stdout},
		showDate:   true,
		showFile:   true,
		dateFormat: time.RFC3339,
	}
}

// AddOutput adds an output writer to the logger
func (l *Logger) AddOutput(w io.Writer) {
	l.outputs = append(l.outputs, w)
}

// SetDateFormat sets the date format
func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

// SetShowDate sets whether to show the date
func (l *Logger) SetShowDate(show bool) {
	l.showDate = show
}

// SetShowFile sets whether to show the file and line number
func (l *Logger) SetShowFile(show bool) {
	l.showFile = show
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.logLevel = level
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.log("DEBUG", format, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.logLevel <= InfoLevel {
		l.log("INFO", format, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.logLevel <= WarnLevel {
		l.log("WARN", format, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.logLevel <= ErrorLevel {
		l.log("ERROR", format, args...)
	}
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.logLevel <= FatalLevel {
		l.log("FATAL", format, args...)
		os.Exit(1)
	}
}

// log logs a message with the given level
func (l *Logger) log(level, format string, args ...interface{}) {
	var sb strings.Builder

	// Add date
	if l.showDate {
		sb.WriteString(time.Now().Format(l.dateFormat))
		sb.WriteString(" ")
	}

	// Add level
	sb.WriteString("[")
	sb.WriteString(level)
	sb.WriteString("] ")

	// Add file and line
	if l.showFile {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			file = filepath.Base(file)
			sb.WriteString(fmt.Sprintf("%s:%d ", file, line))
		}
	}

	// Add message
	sb.WriteString(fmt.Sprintf(format, args...))

	// Write to all outputs
	logLine := sb.String()
	for _, w := range l.outputs {
		fmt.Fprintln(w, logLine)
	}
}

// FileLogger creates a logger that writes to a file
func FileLogger(filename string, level LogLevel) (*Logger, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open log file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create logger
	logger := NewLogger(level)
	logger.outputs = []io.Writer{file}
	return logger, nil
}

// ConsoleLogger creates a logger that writes to the console
func ConsoleLogger(level LogLevel) *Logger {
	return NewLogger(level)
}

// DualLogger creates a logger that writes to both a file and the console
func DualLogger(filename string, level LogLevel) (*Logger, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open log file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create logger
	logger := NewLogger(level)
	logger.outputs = []io.Writer{os.Stdout, file}
	return logger, nil
}

// LogFunc is a function that logs a message
type LogFunc func(format string, args ...interface{})

// TimedOperation logs the duration of an operation
func TimedOperation(logFunc LogFunc, operation string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		logFunc("%s took %s", operation, duration)
	}
}

// SetupDefaultLogger sets up a default logger for the standard log package
func SetupDefaultLogger(filename string, level LogLevel) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open log file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Set up log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	return nil
}

// LevelToString converts a LogLevel to a string
func LevelToString(level LogLevel) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// StringToLevel converts a string to a LogLevel
func StringToLevel(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	case "WARN":
		return WarnLevel
	case "ERROR":
		return ErrorLevel
	case "FATAL":
		return FatalLevel
	default:
		return InfoLevel
	}
} 