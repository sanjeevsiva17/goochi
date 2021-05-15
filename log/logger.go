package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type logger struct {
	out   io.Writer
	level level
	// App Specific Data for the logger
	isTerminal bool
}

// log does the actual logging. This function creates the logEntry message and outputs it in color format
// in terminal context and gives out json in non terminal context. Also, sends to echo if client is present.
func (l *logger) log(level level, args string) {
	if l.level < level {
		return
	}

	e := &logEntry{
		Level:   level,
		Time:    time.Now(),
		Message: args,
		System:  fetchSystemStats(),
	}

	if l.isTerminal {
		fmt.Fprint(l.out, e.TerminalOutput())
	} else {
		_ = json.NewEncoder(l.out).Encode(e)
	}
}

func isJSON(s string) (ok bool, hashmap map[string]interface{}) {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil, js
}

func (l *logger) Log(args ...interface{}) {
	l.log(Info, fmt.Sprint(args...))
}

func (l *logger) Logf(format string, args ...interface{}) {
	l.log(Info, fmt.Sprintf(format, args...))
}

func (l *logger) Info(args ...interface{}) {
	l.log(Info, fmt.Sprint(args...))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log(Info, fmt.Sprintf(format, args...))
}

func (l *logger) Debug(args ...interface{}) {
	l.log(Debug, fmt.Sprint(args...))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log(Debug, fmt.Sprintf(format, args...))
}

func (l *logger) Warn(args ...interface{}) {
	l.log(Warn, fmt.Sprint(args...))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log(Warn, fmt.Sprintf(format, args...))
}

func (l *logger) Error(args ...interface{}) {
	l.log(Error, fmt.Sprint(args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log(Error, fmt.Sprintf(format, args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.log(Fatal, fmt.Sprint(args...))
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log(Fatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func NewLogger(level string) Logger {
	l := &logger{
		out: os.Stdout,
	}

	switch strings.ToUpper(level) {
	case "INFO":
		l.level = Info
	case "WARN":
		l.level = Warn
	case "DEBUG":
		l.level = Debug
	case "FATAL":
		l.level = Fatal
	case "ERROR":
		l.level = Error
	}

	// Set terminal to ensure proper output format.
	l.isTerminal = checkIfTerminal(l.out)

	return l
}
