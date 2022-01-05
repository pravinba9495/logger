package logger

var (
	// Default log Level
	LOG_LEVEL = "ERROR"
)

// SetLogLevel sets the log level
func SetLogLevel(level string) {
	LOG_LEVEL = level
}

// GetLogLevel returns the set log level
func GetLogLevel() int {
	switch LOG_LEVEL {
	case "TRACE":
		return 0
	case "LOG":
		return 1
	case "WARN":
		return 2
	case "ERROR":
		return 3
	default:
		return 3
	}
}
