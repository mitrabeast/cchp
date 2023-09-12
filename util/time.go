package util

import "time"

func MeasureRuntime(f func()) (time.Duration, uint64) {
	start := time.Now()
	startTicks := CpuTicks()
	f()
	endTicks := CpuTicks()
	return time.Since(start), endTicks - startTicks
}
