package logger

import "strings"

// Level represent the log level.
type Level string

const (
	// DebugLevel has verbose message
	DebugLevel Level = "debug"

	// InfoLevel is default log level
	InfoLevel Level = "info"

	// WarnLevel is for logging messages about possible issues
	WarnLevel Level = "warn"

	// ErrorLevel is for logging errors
	ErrorLevel Level = "error"

	// FatalLevel is for logging fatal messages. The system shutdown after logging the message.
	FatalLevel Level = "fatal"
)

var log Logger

func init() {
	defaultConf := Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: false,
		ConsoleLevel:      InfoLevel,
		EnableFile:        false,
	}
	log = newZapLogger(defaultConf)
}

// Logger is our contract for the logger
type Logger interface {
	Debug(format string, args ...interface{})

	Info(format string, args ...interface{})

	Warn(format string, args ...interface{})

	Error(format string, args ...interface{})

	Fatal(format string, args ...interface{})

	Panic(format string, args ...interface{})
}

// Configuration contains all parameters for the logger.
type Configuration struct {
	// Console
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      Level

	// Files
	EnableFile     bool
	FileJSONFormat bool
	FileLevel      Level
	FileLocation   string
}

func GetLevel(l string) Level {
	switch strings.ToLower(l) {
	case "debug":
		return DebugLevel

	case "info":
		return InfoLevel

	case "warn":
		return WarnLevel

	case "error":
		return ErrorLevel

	case "fatal":
		return FatalLevel

	default:
		return InfoLevel
	}
}

// NewLogger returns an instance of logger
func NewLogger(config Configuration) {
	log = newZapLogger(config)
}

// Info logs a message at info level. The message includes any fields passed.
func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

// Warn logs a message at warn level. The message includes any fields passed.
func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

// Debug logs a message at debug level. The message includes any fields passed.
func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

// Error logs a message at error level. The message includes any fields passed.
func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

// Fatal logs a message at fatal level. The message includes any fields passed.
func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}

// Panic logs a message at panic level. The message includes any fields passed.
func Panic(format string, args ...interface{}) {
	log.Panic(format, args...)
}
