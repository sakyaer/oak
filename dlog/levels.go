//+build !nolog

package dlog

import "fmt"

var (
	debugLevel = ERROR
)

// LogLevel represents the levels a debug message can have
type LogLevel int

// Logging levels
const (
	NONE LogLevel = iota
	ERROR
	WARN
	INFO
	VERBOSE
)

// GetLogLevel returns the current log level, i.e WARN or INFO...
func GetLogLevel() LogLevel {
	return debugLevel
}

// SetDebugLevel sets what message levels of debug
// will be printed.
func SetDebugLevel(dL LogLevel) {
	if dL < NONE || dL > VERBOSE {
		Warn("Unknown debug level: ", dL)
		debugLevel = NONE
	} else {
		debugLevel = dL
	}
}

// Error will write a dlog if the debug level is not NONE
func Error(in ...interface{}) {
	if debugLevel > NONE {
		dLog(true, true, in)
	}
}

// Warn will write a dLog if the debug level is higher than ERROR
func Warn(in ...interface{}) {
	if debugLevel > ERROR {
		dLog(true, true, in)
	}
}

// Info will write a dLog if the debug level is higher than WARN
func Info(in ...interface{}) {
	if debugLevel > WARN {
		dLog(true, false, in)
	}
}

// Verb will write a dLog if the debug level is higher than INFO
func Verb(in ...interface{}) {
	if debugLevel > INFO {
		dLog(true, false, in)
	}
}

// SetStringDebugLevel parses the input string as one of the debug levels
func SetStringDebugLevel(debugL string) {

	var dLevel LogLevel
	switch debugL {
	case "INFO":
		dLevel = INFO
	case "VERBOSE":
		dLevel = VERBOSE
	case "ERROR":
		dLevel = ERROR
	case "WARN":
		dLevel = WARN
	case "NONE":
		dLevel = NONE
	default:
		dLevel = ERROR
		fmt.Println("setting dlog level to \"", debugL, "\" failed, it is now set to ERROR")
	}

	SetDebugLevel(dLevel)
}
