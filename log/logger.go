package log

import (
	"os"
)

func newLogger() *logger {
	l := &logger{
		out: os.Stdout,
	}

	// Set terminal to ensure proper output format.
	l.isTerminal = checkIfTerminal(l.out)

	return l
}

func NewLogger() Logger {
	return newLogger()
}
