package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogLevel(t *testing.T) {
	testcases := []struct {
		level  level
		output string
	}{
		{Warn, "DEBUG"}, // when log level is set to WARN, DEBUG log must not be logged
		{Fatal, "WARN"}, // when log level is set to FATAL, WARN or DEBUG log must not be logged
	}

	for i, v := range testcases {
		b := new(bytes.Buffer)
		l := NewMockLogger(b, v.level)

		l.Warn("hello")
		l.Warnf("%d", 1)

		l.Debug("debug")
		l.Debugf("%s", v.level)

		if strings.Contains(b.String(), v.output) {
			t.Errorf("[TESTCASE%d]failed.expected %v\tgot %v\n", i+1, b.String(), v.output)
		}
	}
}

func TestLoggingLevels(t *testing.T) {
	b := new(bytes.Buffer)
	l := NewMockLogger(b, Debug)

	{
		b.Reset()

		l.Info("hi")

		if !strings.Contains(b.String(), "hi") {
			t.Errorf("FAILED INFO, got: %v", b.String())
		}
	}

	{
		b.Reset()

		l.Warn("WARN")

		if !strings.Contains(b.String(), "WARN") {
			t.Errorf("FAILED WARN, got: %v", b.String())
		}
	}

	{
		b.Reset()

		l.Debug("DEBUG")

		if !strings.Contains(b.String(), "DEBUG") {
			t.Errorf("FAILED DEBUG, got: %v", b.String())
		}
	}

	{
		b.Reset()

		l.Error("ERROR")

		if !strings.Contains(b.String(), "ERROR") {
			t.Errorf("FAILED ERROR, got: %v", b.String())
		}
	}

}
