package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestLogLevel(t *testing.T) {
	testcases := []struct {
		level  string
		output string
	}{
		{"warn", "DEBUG"}, // when log level is set to WARN, DEBUG log must not be logged
		{"fatal", "WARN"}, // when log level is set to FATAL, WARN or DEBUG log must not be logged
	}

	level := os.Getenv("LOG_LEVEL")

	defer os.Setenv("LOG_LEVEL", level)

	for i, v := range testcases {
		_ = os.Setenv("LOG_LEVEL", v.level)

		b := new(bytes.Buffer)
		l := NewMockLogger(b)

		l.Warn("hello")
		l.Warnf("%d", 1)

		l.Debug("debug")
		l.Debugf("%s", v.level)

		if strings.Contains(v.output, b.String()) {
			t.Errorf("[TESTCASE%d]failed.expected %v\tgot %v\n", i+1, b.String(), v.output)
		}
	}
}
