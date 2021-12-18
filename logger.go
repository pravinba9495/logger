package logger

import (
	"fmt"
	"log"
	"runtime"

	"github.com/fatih/color"
)

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
