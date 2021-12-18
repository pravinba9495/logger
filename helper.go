package logger

import "os"

// SetLogLevel sets the log level
func SetLogLevel(level string) {
	os.Setenv("LOG_LEVEL", level)
}

// GetLogLevel returns the set log level
func GetLogLevel() int {
	level, bool := os.LookupEnv("LOG_LEVEL")
	if bool {
		switch level {
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
	} else {
		if err := os.Setenv("LOG_LEVEL", "ERROR"); err != nil {
			return 3
		}
		return GetLogLevel()
	}
}
