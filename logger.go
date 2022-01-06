package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/fatih/color"
)

var (
	currentLogLevel LogLevel = LevelError
	logFilePath     string   = "logfile.txt"
)

// Initialize logger
func Init(opts *LoggerOptions) (func(), error) {
	if opts != nil {
		if opts.LogLevel != "" {
			currentLogLevel = opts.LogLevel
		}
		if opts.LogFilePath != "" {
			logFilePath = opts.LogFilePath
		}
	}
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return func() {}, err
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	return func() {
		defer f.Close()
	}, nil
}

// GetLogLevel returns the current log level
func GetLogLevel() int {
	switch currentLogLevel {
	case LevelTrace:
		return 0
	case LevelLog:
		return 1
	case LevelWarn:
		return 2
	case LevelError:
		return 3
	default:
		return 3
	}
}

// Trace prints logs of type "TRACE"
func Trace(s string) {
	if GetLogLevel() > 0 {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	str := color.CyanString("%s", s)
	log.Print("[" + filename + ":" + fmt.Sprint(line) + "] " + str)
}

// Print prints logs of type "LOG"
func Print(s string) {
	if GetLogLevel() > 1 {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	str := color.WhiteString("%s", s)
	log.Print("[" + filename + ":" + fmt.Sprint(line) + "] " + str)
}

// Print prints logs of type "SUCCESS"
func Success(s string) {
	_, filename, line, _ := runtime.Caller(1)
	str := color.GreenString("%s", s)
	log.Print("[" + filename + ":" + fmt.Sprint(line) + "] " + str)
}

// Warn prints logs of type "WARN"
func Warn(s string) {
	if GetLogLevel() > 2 {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	str := color.YellowString("%s", s)
	log.Print("[" + filename + ":" + fmt.Sprint(line) + "] " + str)
}

// Error prints logs of type "ERROR"
func Error(s string) {
	if GetLogLevel() > 3 {
		return
	}
	_, filename, line, _ := runtime.Caller(1)
	str := color.RedString("%s", s)
	log.Print("[" + filename + ":" + fmt.Sprint(line) + "] " + str)
}
