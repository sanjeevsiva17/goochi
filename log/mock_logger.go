package log

import (
	"io"
)

func NewMockLogger(output io.Writer, level level) Logger {
	return &logger{
		out:   output,
		level: level,
	}
}
