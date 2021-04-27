package log

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEntryFromStrings(t *testing.T) {
	testcases := []struct {
		args          []string
		expectedEntry *logEntry
	}{
		{[]string{}, &logEntry{Data: map[string]interface{}{}, Message: []string{}}},
		{[]string{"hello logging"}, &logEntry{Data: map[string]interface{}{},
			Message: "hello logging"}},
		{[]string{"hello", "logging"}, &logEntry{Data: map[string]interface{}{},
			Message: []string{"hello", "logging"}}},
	}

	for i, v := range testcases {
		e := dataFromStrings(v.args...)
		if !reflect.DeepEqual(e.Data, v.expectedEntry.Data) || !reflect.DeepEqual(e.Message, v.expectedEntry.Message) {
			t.Errorf("[TESTCASE%d]Failed.Expected Data:%v Message %v\nGot Data:%v Message %v\n", i+1,
				v.expectedEntry.Data, v.expectedEntry.Message, e.Data, e.Message)
		}
	}
}

func TestEntryFromStringForJSON(t *testing.T) {
	args := `{"message":"hello","responseCode":200}`

	expectedEntry := logEntry{Data: map[string]interface{}{"responseCode": 200.00}, Message: "hello"}

	e := dataFromStrings(args)
	if !reflect.DeepEqual(e.Data, expectedEntry.Data) {
		t.Errorf("expected data %v\tgot%v\n", expectedEntry.Data, e.Data)
	}

	if !reflect.DeepEqual(e.Message, expectedEntry.Message) {
		t.Errorf("expected message %v\tgot%v\n", expectedEntry.Message, e.Message)
	}
}

func TestEntry_TerminalOutput(t *testing.T) {
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
			"INFO\u001B[0m [00:00:00]  hello\n%15s: %s")},
		// data with message
		{logEntry{Level: Warn, Message: "hello", Data: map[string]interface{}{"name": "ZopSmart"}},
			"WARN\u001B[0m [00:00:00]  hello"},
		// test data
		{logEntry{Level: Debug, Data: map[string]interface{}{"method": "get", "duration": 10000.0, "uri": "i"}},
			fmt.Sprintf("DEBU\u001B[0m [00:00:00] \u001B[37m %s\u001B[0m %v - %.2fm", "get", "i", 10.0)},
		// app data
		{logEntry{Level: Info, Message: "test"}, fmt.Sprintf(
			"INFO\u001B[0m [00:00:00]  test\n%15s: %v", "a", "b")},
	}

	for i, v := range testcases {
		output := v.input.TerminalOutput()
		if !strings.Contains(output, v.output) {
			t.Errorf("[TESTCASE%d]got %v\texpected %v\n", i+1, output, v.output)
		}
	}
}
