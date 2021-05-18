package log

import (
	"fmt"
	"time"
)

type logEntry struct {
	Level   level                  `json:"level"`
	Time    time.Time              `json:"time"`
	Message interface{}            `json:"message"`
	System  map[string]interface{} `json:"system"`
}

func (l *logEntry) TerminalOutput() string {
	levelColor := l.Level.colorCode()

	s := fmt.Sprintf("\u001B[%dm%s\u001B[0m [%s] ", levelColor, l.Level.String()[0:4], l.Time.Format("15:04:05"))

	s += fmt.Sprintf(" %v", l.Message)

	if l.System != nil {
		s += fmt.Sprintf("\n%15s \u001B[%dm (Memory: %v GoRoutines: %v) \u001B[0m", "", 37, l.System["alloc"], l.System["goRoutines"])
	}

	return s + "\n"
}
