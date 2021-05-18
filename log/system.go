package log

import (
	"os"
	"runtime"
)

func fetchSystemStats() map[string]interface{} {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	stats := make(map[string]interface{})
	stats["alloc"] = m.Alloc
	stats["totalAlloc"] = m.TotalAlloc
	stats["sys"] = m.Sys
	stats["numGC"] = m.NumGC
	stats["goRoutines"] = runtime.NumGoroutine()
	stats["host"], _ = os.Hostname()
	stats["dataCenter"] = os.Getenv("DATA_CENTER")

	if stats["dataCenter"] == "" {
		stats["dataCenter"] = "NOTSET"
	}

	return stats
}
