package logger

// LogLevel defines the log level of the logger instance
type LogLevel string

// LoggerOptions to pass to the logger instance
type LoggerOptions struct {
	LogLevel       LogLevel
	LogFilePath    string
	LogFileMaxSize int64
}
