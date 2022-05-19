package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var (
	currentLogLevel LogLevel = LevelError
	logFilePath     string   = "logfile.txt"
	maxSizeInBytes  int64    = 5 * 1024 * 1024
)

// Init initializes a logger
func Init(opts *LoggerOptions) (func(), error) {
	if opts != nil {
		if opts.LogLevel != "" {
			currentLogLevel = opts.LogLevel
		}
		if opts.LogFilePath != "" {
			logFilePath = opts.LogFilePath
		}
		if opts.LogFileMaxSize > 0 {
			maxSizeInBytes = opts.LogFileMaxSize
		}
	}
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return func() {}, err
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	go func() {
		for {
			size, err := GetFileSize()
			if err != nil {
				panic(err)
			}
			if size >= maxSizeInBytes {
				source, err := os.Open(logFilePath)
				if err != nil {
					panic(err)
				}
				destination, err := os.Create(logFilePath + "-" + time.Now().Format("20060102150405"))
				if err != nil {
					panic(err)
				}
				_, err = io.Copy(destination, source)
				if err != nil {
					panic(err)
				}
				source.Close()
				destination.Close()
				if err := os.Truncate(logFilePath, 0); err != nil {
					panic(err)
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()

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

// Success prints logs of type "SUCCESS"
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

// GetFileSize returns the log file size
func GetFileSize() (int64, error) {
	info, err := os.Stat(logFilePath)
	if err != nil {
		return 0, err
	}
	if !info.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", logFilePath)
	}
	bytes := info.Size()
	return bytes, nil
}
