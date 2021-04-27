package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type logger struct {
	out io.Writer

	// App Specific Data for the logger
	isTerminal bool
}

// log does the actual logging. This function creates the logEntry message and outputs it in color format
// in terminal context and gives out json in non terminal context. Also, sends to echo if client is present.
func (k *logger) log(level level, args ...string) {
	//if rls.level < level {
	//	return // No need to do anything if we are not going to log it.
	//}

	e := dataFromStrings(args...)
	e.Level = level
	e.System = fetchSystemStats()

	if k.isTerminal {
		fmt.Fprint(k.out, e.TerminalOutput())
	} else {
		_ = json.NewEncoder(k.out).Encode(e)
	}
}

func isJSON(s string) (ok bool, hashmap map[string]interface{}) {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil, js
}

func (k *logger) Log(args ...interface{}) {
	k.log(Info, fmt.Sprint(args...))
}

func (k *logger) Logf(format string, args ...interface{}) {
	k.log(Info, fmt.Sprintf(format, args...))
}

func (k *logger) Info(args ...interface{}) {
	k.log(Info, fmt.Sprint(args...))
}

func (k *logger) Infof(format string, args ...interface{}) {
	k.log(Info, fmt.Sprintf(format, args...))
}

func (k *logger) Debug(args ...interface{}) {
	k.log(Debug, fmt.Sprint(args...))
}

func (k *logger) Debugf(format string, args ...interface{}) {
	k.log(Debug, fmt.Sprintf(format, args...))
}

func (k *logger) Warn(args ...interface{}) {
	k.log(Warn, fmt.Sprint(args...))
}

func (k *logger) Warnf(format string, args ...interface{}) {
	k.log(Warn, fmt.Sprintf(format, args...))
}

func (k *logger) Error(args ...interface{}) {
	k.log(Error, fmt.Sprint(args...))
}

func (k *logger) Errorf(format string, args ...interface{}) {
	k.log(Error, fmt.Sprintf(format, args...))
}

func (k *logger) Fatal(args ...interface{}) {
	k.log(Fatal, fmt.Sprint(args...))
	os.Exit(1)
}

func (k *logger) Fatalf(format string, args ...interface{}) {
	k.log(Fatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}
