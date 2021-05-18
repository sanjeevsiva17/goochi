package log

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestEntry_TerminalOutput(t *testing.T) {
	sysStats := fetchSystemStats()
	now := time.Now()
	formattedNow := now.Format("15:04:05")
	testcases := []struct {
		input  logEntry
		output string
	}{
		// fatal error checking if msg and level is logged
		{logEntry{Level: Fatal, Message: "fatal error", Time: now},
			"FATA\u001B[0m [" + formattedNow + "]  fatal error"},
		// correlationId
		{logEntry{Level: Info, Message: "hello"}, fmt.Sprintf(
			"INFO\u001B[0m [00:00:00]  hello\n")},
		// data with message
		{logEntry{Level: Warn, Message: "hello"},
			"WARN\u001B[0m [00:00:00]  hello"},
		// system
		{logEntry{Level: Info, Message: "hi", System: sysStats},
			fmt.Sprintf("\u001B[36mINFO\u001B[0m [00:00:00]  hi\n                \u001B[37m (Memory: %v GoRoutines: %v) \u001B[0m\n", sysStats["alloc"], sysStats["goRoutines"])},
	}

	for i, v := range testcases {
		output := v.input.TerminalOutput()
		if !strings.Contains(output, v.output) {
			t.Errorf("[TESTCASE%d]\ngot %v\nexp %v\n", i+1, output, v.output)
		}
	}
}
