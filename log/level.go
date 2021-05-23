package log

import (
	"strings"
)

type level int

const (
	Fatal level = iota + 1
	Error
	Warn
	Info
	Debug
)

func (l level) String() string {
	switch l {
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Debug:
		return "DEBUG"
	default:
		return "INFO"
	}
}

const (
	redColor    = 31
	yellowColor = 33
	blueColor   = 36
	normalColor = 37
)

// colorCode returns the color to be used for the formatting at terminal
func (l level) colorCode() int {
	switch l {
	case Error, Fatal:
		return redColor
	case Warn:
		return yellowColor
	case Info:
		return blueColor
	default:
		return normalColor
	}
}

func getLevel(level string) level {
	switch strings.ToUpper(level) {
	case "INFO":
		return Info
	case "WARN":
		return Warn
	case "FATAL":
		return Fatal
	case "DEBUG":
		return Debug
	case "ERROR":
		return Error
	default:
		return Info
	}
}
