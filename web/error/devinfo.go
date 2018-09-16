package error

import (
	"runtime"
	"runtime/debug"
)

type devInfo struct {
	stacktrace []byte
	file       string
	line       int
}

//go:noinline
func getDevInfo(offset int) *devInfo {
	stacktrace := debug.Stack()
	_, file, line, ok := runtime.Caller(offset + 1)
	if !ok {
		file = "???"
		line = 0
	}

	return &devInfo{stacktrace, file, line}
}
