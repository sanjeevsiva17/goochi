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
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (l *logEntry) TerminalOutput() string {
	levelColor := l.Level.colorCode()

	s := fmt.Sprintf("\u001B[%dm%s\u001B[0m [%s] ", levelColor, l.Level.String()[0:4], l.Time.Format("15:04:05"))

	if len(l.Data) > 0 {
		if l.Message != nil { // http client sends message on error
			s += fmt.Sprintf(" %s", l.Message)
		} else {
			//nolint:mnd // 1000 is used to convert the microsecond to milliseconds
			s += fmt.Sprintf("\u001B[%dm %s\u001B[0m %v - %.2fms", 37, l.Data["method"], l.Data["uri"], l.Data["duration"].(float64)/1000)
		}
	} else {
		s += fmt.Sprintf(" %v", l.Message)
	}

	if l.System != nil {
		s += fmt.Sprintf("\n%15s \u001B[%dm (Memory: %v GoRoutines: %v) \u001B[0m", "", 37, l.System["alloc"], l.System["goRoutines"])
	}

	return s + "\n"
}

// dataFromStrings takes multiple strings and creates a log logEntry from it.
// It is written separately so different use-cases can be tested safely in isolation.
// nolint
func dataFromStrings(args ...string) *logEntry {
	l := logEntry{
		Time: time.Now(),
		Data: make(map[string]interface{}),
	}

	// No need for array if size is only 1
	if len(args) == 1 {
		j, valMap := isJSON(args[0])
		if j {
			l.Data = valMap

			if m, ok := valMap["message"]; ok {
				l.Message = m.(string)

				delete(valMap, "message")
			}
		} else {
			l.Message = args[0]
		}
	} else {
		l.Message = args
	}

	return &l
}
