package log

import "fmt"

// Level defines a list of log priorities
type Level int

const (
	// INFO ...
	INFO Level = 0
	// DEBUG ...
	DEBUG Level = 1
	// WARN ...
	WARN Level = 2
	// ERROR ...
	ERROR Level = 3
)

func (l Level) toString() string {
	switch l {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}

	return "Not supported"
}

// DefaultLevel ...
const DefaultLevel Level = INFO

var logLevel Level

// SetLogLevel ...
func SetLogLevel(level Level) {
	logLevel = level
}

// Info ...
func Info(f string, msg ...interface{}) {
	log(INFO, f, msg...)
}

// Debug ...
func Debug(f string, msg ...interface{}) {
	log(DEBUG, f, msg...)
}

// Warn ...
func Warn(f string, msg ...interface{}) {
	log(WARN, f, msg...)
}

// Error ...
func Error(f string, msg ...interface{}) {
	log(ERROR, f, msg...)
}

func log(l Level, f string, msg ...interface{}) {
	if logLevel < l {
		return
	}

	internalLog(format(f), message(l, msg...)...)
}

func internalLog(f string, msg ...interface{}) {
	fmt.Printf(f, msg...)
}

func format(f string) string {
	return fmt.Sprintf("%s%s%s\n", "%s", "%s", f)
}

func message(level Level, msg ...interface{}) []interface{} {
	var msgs []interface{}
	msgs = append(msgs, level.toString(), ": ")
	msgs = append(msgs, msg...)

	return msgs
}
