package log

import (
	"io"
)

func NewMockLogger(output io.Writer) Logger {
	return &logger{
		out: output,
	}
}
